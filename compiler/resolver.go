package compiler

import (
	"fmt"
	"path"
	"strings"
)

func (c *Compiler) resolve() error {
	resolvers := []func() error{
		c.resolveImports, // Must be done first
		c.resolveFieldNumbers,
	}

	for _, resolver := range resolvers {
		if err := resolver(); err != nil {
			return err
		}
	}

	return nil
}

// resolveImports resolves all imports in the scripts. This is one of the most important steps and should be done first.
// All scripts in the same package get their siblings imported.
func (c *Compiler) resolveImports() error {
	return c.resolveImportsForPackage(c.srcPackage)
}

func (c *Compiler) resolveImportsForPackage(pkg *ScriptPackage) error {
	var addImportToScript func(sc *Script, imp *ImportStatement) error

	makeAccessiable := func(curr *Script, target *Script) {
		curr.AccessibleScripts = append(curr.AccessibleScripts, target)

		if target.AreImportsPrivate {
			return // no need to go further
		}

		for _, imp := range target.Imports {
			if imp.Private {
				continue
			}

			addImportToScript(curr, imp)
		}
	}

	addImportToScript = func(sc *Script, imp *ImportStatement) error {
		// if path has .proto extension, skip
		if strings.HasSuffix(imp.Path, ".proto") {
			return nil
		}

		// first, convert import path to fullpath, @ is alias to src folder
		fullpath := ""
		if strings.HasPrefix(imp.Path, "@") {
			fullpath = path.Join(c.srcDir, imp.Path[1:])
		} else if strings.HasPrefix(imp.Path, ".") {
			// its a relative path, resolve it relative to the current pkg path
			fullpath = path.Join(pkg.Path, imp.Path)
		} else {
			// its an absolute path
			fullpath = imp.Path
		}

		// if path has .plsd extension, try to find the
		if strings.HasSuffix(fullpath, ".plsd") {
			impScript := c.findScriptByPath(fullpath)
			if impScript == nil {
				return fmt.Errorf("imported Script '%s' not found", fullpath)
			}

			makeAccessiable(sc, impScript)

			return nil
		}

		// else, its probably a package, import all scripts in the package
		impPkg := c.findPackage(fullpath)
		if impPkg == nil {
			return fmt.Errorf("imported package '%s' not found", fullpath)
		}

		for _, impScript := range impPkg.Scripts {
			makeAccessiable(sc, impScript)
		}

		return nil
	}

	for _, sc := range pkg.Scripts {
		sc.AccessibleScripts = []*Script{}
		for _, osc := range pkg.Scripts { // add all siblings to the context
			if osc == sc {
				continue
			}
			sc.AccessibleScripts = append(sc.AccessibleScripts, osc)
		}

		for _, imp := range sc.Imports {
			if err := addImportToScript(sc, imp); err != nil {
				return err
			}
		}
	}

	for _, subPkg := range pkg.Packages {
		if err := c.resolveImportsForPackage(subPkg); err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) resolveFieldNumbers() error {
	for _, sc := range c.scripts {
		for _, msg := range sc.Messages {
			var assignNumbers func(*Message) error
			assignNumbers = func(msg *Message) error {
				usedNumbers := map[int]bool{}
				biggestNumber := 1

				// first iteration: find all used numbers and the biggest assigned number to start from
				for _, field := range msg.Fields {
					if field.Order <= 0 {
						assignedNumber := getAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name)
						if assignedNumber != nil {
							if usedNumbers[*assignedNumber] {
								deleteAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name)
							} else {
								if *assignedNumber > biggestNumber {
									biggestNumber = *assignedNumber + 1
								}

								usedNumbers[*assignedNumber] = true
								field.Order = *assignedNumber
							}
						}

						continue
					}

					if usedNumbers[field.Order] {
						return fmt.Errorf("field number %d is already used in Message '%s'", field.Order, msg.Name)
					}

					usedNumbers[field.Order] = true

					if field.Order > biggestNumber {
						biggestNumber = field.Order + 1
					}

					assignedNumber := getAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name)
					if assignedNumber != nil {
						if usedNumbers[*assignedNumber] {
							deleteAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name)
						} else {
							if *assignedNumber != field.Order {
								if field.Overwrite {
									setAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name, field.Order)
								} else {
									field.Order = *assignedNumber
								}
							}
						}
					}
				}

				// second iteration: assign numbers to fields that have not been assigned yet, and save them
				currentNumber := biggestNumber
				getNextNumber := func() int {
					for {
						if !usedNumbers[currentNumber] && !isReservedFieldNumber(currentNumber) {
							usedNumbers[currentNumber] = true
							return currentNumber
						}
						currentNumber++
					}
				}

				for _, field := range msg.Fields {
					if field.Order <= 0 {
						field.Order = getNextNumber()

						if c.config.OrderPersist != nil && *c.config.OrderPersist {
							setAssignedNumber(c.mapping, sc.Path, msg.Name, field.Name, field.Order)
						}
					}
				}

				if msg.Children != nil {
					for _, child := range msg.Children {
						if err := assignNumbers(child); err != nil {
							return err
						}
					}
				}

				return nil
			}

			if err := assignNumbers(msg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Compiler) findPackage(path string) *ScriptPackage {
	return c.findPackageInPackage(path, c.srcPackage)
}

func (c *Compiler) findPackageInPackage(path string, pkg *ScriptPackage) *ScriptPackage {
	if path == "" {
		return pkg
	}

	for p, sp := range pkg.Packages {
		if p == path {
			return sp
		}

		return c.findPackageInPackage(path, sp)
	}

	return nil
}

func (c *Compiler) findScriptByPath(path string) *Script {
	return c.findScriptByPathInPackage(path, c.srcPackage)
}

func (c *Compiler) findScriptByPathInPackage(path string, pkg *ScriptPackage) *Script {
	for p, sc := range pkg.Scripts {
		if p == path {
			return sc
		}
	}

	for _, sp := range pkg.Packages {
		if sc := c.findScriptByPathInPackage(path, sp); sc != nil {
			return sc
		}
	}

	return nil
}

// find finds a Message or Enum by its identifier. The identifier can be concatenated with dots to represent a nested Message.
// A Message can be in one of two nested places: another package or another Message. messages have higher precedence than packages.
func (p *Script) find(identifier string) *findResult {
	parts := strings.Split(identifier, ".")
	if len(parts) == 1 {
		if msg, ok := p.Messages[identifier]; ok {
			return &findResult{Message: msg}
		}

		if enm, ok := p.Enums[identifier]; ok {
			return &findResult{Enum: enm}
		}

		for _, sc := range p.AccessibleScripts {
			if sc.Package == p.Package {
				if msg, ok := sc.Messages[identifier]; ok {
					return &findResult{Message: msg}
				}

				if enm, ok := sc.Enums[identifier]; ok {
					return &findResult{Enum: enm}
				}
			}
		}

		return nil
	}

	if len(parts) == 2 {
		for _, sc := range p.AccessibleScripts {
			if sc.Package != p.Package {
				if sc.Package.Name == parts[0] {
					if msg, ok := sc.Messages[parts[1]]; ok {
						return &findResult{Message: msg}
					}

					if enm, ok := sc.Enums[parts[1]]; ok {
						return &findResult{Enum: enm}
					}
				}
			}
		}
	}

	var findNestedMessage func(msg *Message, parts []string) *findResult
	findNestedMessage = func(msg *Message, parts []string) *findResult {
		if len(parts) <= 0 {
			return &findResult{Message: msg}
		}

		if nestedMsg, ok := msg.Children[parts[0]]; ok {
			return findNestedMessage(nestedMsg, parts[1:])
		}

		return nil
	}

	msg, ok := p.Messages[parts[0]]
	if !ok {
		return nil
	}

	return findNestedMessage(msg, parts[1:])
}

func (sc *Script) isTypeDeclared(t string) bool {
	if _, ok := sc.DeclaredTypes[t]; ok {
		return true
	}

	for _, asc := range sc.AccessibleScripts {
		if asc.Package == sc.Package {
			if _, ok := asc.DeclaredTypes[t]; ok {
				return true
			}
		}
	}

	return false
}

type findResult struct {
	Message *Message
	Enum    *Enum
}

func (p *Script) resolveDataTypeAlias(possibleAlias string) string {
	if possibleAlias == "int" {
		return "int32"
	}

	if possibleAlias == "uint" {
		return "uint32"
	}

	if possibleAlias == "sint" {
		return "sint32"
	}

	if possibleAlias == "long" {
		return "int64"
	}

	if possibleAlias == "ulong" {
		return "uint64"
	}

	alias, ok := p.TypeAliases[possibleAlias]
	if !ok {
		return possibleAlias
	}

	return alias
}

func isNativeType(identifier string) bool {
	switch identifier {
	case "double", "float", "int32", "int64", "uint32", "uint64", "sint32", "sint64", "fixed32", "fixed64", "sfixed32", "sfixed64", "bool", "string", "bytes":
		return true
	}

	return false
}

func isReservedFieldNumber(number int) bool {
	return number >= 19000 && number <= 19999
}
