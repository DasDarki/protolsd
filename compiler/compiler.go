package compiler

import (
	"encoding/json"
	"os"
	"path"
	"protolsd/parser"
	"protolsd/protobuf"
	"protolsd/util"

	"github.com/antlr4-go/antlr/v4"
)

type Compiler struct {
	config     *Config
	baseDir    string
	srcDir     string
	srcPackage *scriptPackage
	scripts    []*script
	mapping    fieldNumberMapping
	loggger    *util.Logger
	lspMode    bool
}

func NewCompiler(config *Config, baseDir string, lspMode bool, logger *util.Logger) *Compiler {
	c := &Compiler{
		config:  config,
		baseDir: baseDir,
		srcDir:  path.Join(baseDir, config.InputDir),
		srcPackage: &scriptPackage{
			Path:              "",
			Name:              unptr(config.Protobuf.Package),
			AreImportsPrivate: false,
			Scripts:           map[string]*script{},
			Packages:          map[string]*scriptPackage{},
		},
		scripts: []*script{},
		mapping: fieldNumberMapping{},
		lspMode: lspMode,
		loggger: logger,
	}

	if c.config.OrderPersist != nil && *c.config.OrderPersist {
		mapping, err := readFieldNumberMapping(path.Join(baseDir, ".protolsd_persist"))
		if err != nil {
			c.loggger.Err("Failed to read field number mapping: %v", err)
		} else {
			c.mapping = mapping
		}
	}

	return c
}

func (c *Compiler) Compile() error {
	if err := c.analyzeDir(c.srcDir, c.srcPackage); err != nil {
		return err
	}

	if c.loggger.IsDebug() {
		c.Debug()
	}

	if err := c.resolve(); err != nil {
		return err
	}

	if c.lspMode {
		return nil
	}

	inputDir, err := c.generate()
	if err != nil {
		return err
	}

	if c.config.OrderPersist != nil && *c.config.OrderPersist {
		if err := saveFieldNumberMapping(c.mapping, path.Join(c.baseDir, ".protolsd_persist")); err != nil {
			c.loggger.Warn("Failed to save field number mapping: %v", err)
		}
	}

	if c.config.OutputType != nil && *c.config.OutputType == OutputTypeCompiled {
		if c.config.Protobuf.AutoDL != nil && !(*c.config.Protobuf.AutoDL) {
			c.loggger.Crit("Protobuf auto download is disabled, compiling without protoc not supported!")
			return nil
		}

		protocDir, err := protobuf.DownloadProtobuf(*c.config.Protobuf.Version, c.loggger)
		if err != nil {
			return err
		}

		c.loggger.Info("Using protoc from %s", protocDir)

		if *c.config.WithGRPC {
			if err := protobuf.CheckGrpcPrerequisities(); err != nil {
				return err
			}
		}

		for _, target := range c.config.Protobuf.Targets {
			outDirname := target

			if c.config.Protobuf.TargetOptions != nil {
				if opts, ok := c.config.Protobuf.TargetOptions[target]; ok {
					if outDir, ok := opts["output_dir"]; ok {
						outDirname = outDir
					}
				}
			}

			outDir := path.Join(c.baseDir, *c.config.OutputDir, outDirname)

			exe := protobuf.NewExecutor(protobuf.Target(target), inputDir, outDir, protocDir, *c.config.WithGRPC)
			if err := exe.Execute(c.loggger); err != nil {
				return err
			}

			c.loggger.Info("Compiled %s to %s", inputDir, outDir)
		}
	}

	return nil
}

func (c *Compiler) Debug() {
	jsonText, err := json.MarshalIndent(c.srcPackage, "", "  ")
	if err != nil {
		c.loggger.Err("Failed to marshal script to JSON: %v", err)
		return
	}

	println(string(jsonText))
}

func (c *Compiler) analyzeDir(dir string, pkg *scriptPackage) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			childPkg := &scriptPackage{
				Name:              entry.Name(),
				Path:              path.Join(dir, entry.Name()),
				AreImportsPrivate: false,
				Parent:            pkg,
				Scripts:           map[string]*script{},
				Packages:          map[string]*scriptPackage{},
			}

			pkg.Packages[childPkg.Name] = childPkg

			if err := c.analyzeDir(path.Join(dir, entry.Name()), childPkg); err != nil {
				return err
			}
		} else if path.Ext(entry.Name()) == ".plsd" {
			if err := c.anaylzeFile(path.Join(dir, entry.Name()), pkg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Compiler) anaylzeFile(file string, pkg *scriptPackage) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	is := antlr.NewInputStream(string(data))
	lexer := parser.NewProtoLSDLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewProtoLSDParser(stream)

	a := &analyzer{
		script: &script{
			Package:       pkg,
			Path:          file,
			Imports:       []*importStatement{},
			TypeAliases:   map[string]string{},
			DeclaredTypes: map[string]bool{},
			Enums:         map[string]*enum{},
			Messages:      map[string]*message{},
			Services:      map[string]*service{},
		},
		inGlobalScope:      true,
		visitedMessageDefs: map[*parser.MessageDefinitionContext]bool{},
	}

	antlr.ParseTreeWalkerDefault.Walk(a, p.ProtoLSD())

	pkg.Scripts[file] = a.script
	c.scripts = append(c.scripts, a.script)

	return nil
}
