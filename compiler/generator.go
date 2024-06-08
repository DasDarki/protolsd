package compiler

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

func (c *Compiler) generate() (string, error) {
	outdir := ""
	if c.config.OutputDir != nil {
		outdir = path.Join(c.baseDir, *c.config.OutputDir)

		if c.config.OutputType != nil && *c.config.OutputType == OutputTypeCompiled {
			outdir = path.Join(outdir, ".lsd_transpiled")
		}
	}

	if outdir == "" {
		return "", errors.New("output directory not set")
	}

	if err := c.generatePackage(outdir, c.srcPackage); err != nil {
		return "", err
	}

	return outdir, nil
}

func (c *Compiler) generatePackage(dir string, pkg *ScriptPackage) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	var opts map[string]string
	if c.config.Protobuf != nil && c.config.Protobuf.Options != nil {
		opts = c.config.Protobuf.Options
	}

	generateCommons := false
	usedEmptyRequest := false
	usedEmptyResponse := false
	addedImports := map[string]bool{}

	builder := newProtoBuilder()
	builder.AddSyntax("proto3")
	builder.AddPackage(pkg.Name)

	if opts != nil {
		builder.AddOptions(opts)
	}

	for _, sc := range pkg.Scripts {
		if sc.UsedEmptyRequest || sc.UsedEmptyResponse {
			generateCommons = true

			if sc.UsedEmptyRequest {
				usedEmptyRequest = true
			}

			if sc.UsedEmptyResponse {
				usedEmptyResponse = true
			}

			if _, ok := addedImports["commons.proto"]; !ok {
				builder.AddImport("commons.proto")
				addedImports["commons.proto"] = true
			}
		}

		for _, imp := range sc.Imports {
			if strings.HasSuffix(imp.Path, ".proto") {
				if _, ok := addedImports[imp.Path]; !ok {
					builder.AddImport(imp.Path)
					addedImports[imp.Path] = true
				}
			}
		}
	}

	builder.addLine("")

	for _, sc := range pkg.Scripts {
		for _, enm := range sc.Enums {
			values := map[int]string{}
			prefix := ""

			if enm.IsClass {
				prefix = enm.Name + "_"
			}

			for key, val := range enm.Constants {
				values[val] = prefix + key
			}

			builder.AddEnum(enm.Name, values)
		}
	}

	for _, sc := range pkg.Scripts {
		var generateMessage func(*Message, *ProtoBuilder)
		generateMessage = func(m *Message, pb *ProtoBuilder) {
			fields := map[string]ProtoField{}

			for _, f := range m.Fields {
				if !isDataTypeValidInContext(sc, f.DataType) {
					panic(fmt.Errorf("invalid data type %s for %s of %s in context of Script %s", f.DataType.Text, f.Name, m.Name, sc.Path))
				}

				modifiers := []string{}

				if f.DataType.IsOptional {
					modifiers = append(modifiers, "optional")
				}

				if f.DataType.IsArray {
					modifiers = append(modifiers, "repeated")
				}

				fields[f.Name] = ProtoField{
					Type:     f.DataType.Text,
					Number:   f.Order,
					Modifier: strings.Join(modifiers, " "),
				}
			}

			pb.AddMessage(m.Name, fields, func(pb *ProtoBuilder) {
				if m.Children != nil {
					for _, child := range m.Children {
						generateMessage(child, pb)
					}
				}
			})
		}

		for _, msg := range sc.Messages {
			generateMessage(msg, builder)
		}

		for _, s := range sc.Services {
			for _, rpc := range s.Rpcs {
				if rpc.Input != nil && rpc.Input.Message != nil {
					generateMessage(rpc.Input.Message, builder)
				}

				if rpc.Returns != nil && rpc.Returns.Message != nil {
					generateMessage(rpc.Returns.Message, builder)
				}
			}
		}
	}

	for _, sc := range pkg.Scripts {
		for _, s := range sc.Services {
			methods := map[string][2]string{}

			for _, rpc := range s.Rpcs {
				input := ""
				output := ""

				if rpc.Input == nil {
					input = "common.EmptyRequest"
				} else {
					if rpc.Input.Message == nil {
						msgName := sc.ResolveDataTypeAlias(*rpc.Input.Name)

						if sc.IsTypeDeclared(msgName) || isMessageAvailable(sc, msgName) {
							input = msgName
						} else {
							panic(fmt.Errorf("Message %s not found in Script %s", msgName, sc.Path))
						}
					} else {
						input = rpc.Input.Message.Name
					}
				}

				if rpc.Returns == nil {
					output = "common.EmptyResponse"
				} else {
					if rpc.Returns.Message == nil {
						msgName := sc.ResolveDataTypeAlias(*rpc.Returns.Name)

						if sc.IsTypeDeclared(msgName) || isMessageAvailable(sc, msgName) {
							output = msgName
						} else {
							panic(fmt.Errorf("Message %s not found in Script %s", msgName, sc.Path))
						}
					} else {
						output = rpc.Returns.Message.Name
					}
				}

				methods[rpc.Name] = [2]string{input, output}
			}

			builder.AddService(s.Name, methods)
		}
	}

	if generateCommons {
		if err := c.generateCommons(dir, usedEmptyResponse, usedEmptyRequest, opts); err != nil {
			return err
		}
	}

	code := strings.TrimSpace(builder.String())
	if code != "" {
		if err := writeProtoFile(path.Join(dir, pkg.Name+".proto"), code); err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) generateCommons(dir string, res, req bool, opts map[string]string) error {
	builder := newProtoBuilder()
	builder.AddSyntax("proto3")
	builder.AddPackage("common")

	if opts != nil {
		builder.AddOptions(opts)
	}

	if req {
		builder.AddMessage("EmptyRequest", map[string]ProtoField{}, nil)
	}

	if res {
		builder.AddMessage("EmptyResponse", map[string]ProtoField{}, nil)
	}

	code := strings.TrimSpace(builder.String())
	if code != "" {
		if err := writeProtoFile(path.Join(dir, "commons.proto"), code); err != nil {
			return err
		}
	}

	return nil
}

func isDataTypeValidInContext(sc *Script, dt *DataType) bool {
	dt.Text = sc.ResolveDataTypeAlias(dt.Text)
	t := dt.Text

	if IsNativeType(t) {
		return true
	}

	if sc.IsTypeDeclared(t) {
		return true
	}

	result := sc.FindMessageOrEnum(t)
	if result == nil || (result.Enum == nil && result.Message == nil) {
		return false
	}

	return true
}

func isMessageAvailable(sc *Script, name string) bool {
	result := sc.FindMessageOrEnum(name)
	return result != nil && result.Message != nil
}

func writeProtoFile(filename, code string) error {
	return os.WriteFile(filename, []byte(code), 0644)
}
