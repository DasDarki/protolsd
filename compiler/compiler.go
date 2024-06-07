package compiler

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"protolsd/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Compiler struct {
	config  *Config
	scripts map[string]*script
}

func NewCompiler(config *Config) *Compiler {
	return &Compiler{
		config:  config,
		scripts: map[string]*script{},
	}
}

func (c *Compiler) Compile(baseDir string) error {
	srcDir := path.Join(baseDir, c.config.InputDir)

	return c.compileDir(baseDir, srcDir)
}

func (c *Compiler) Debug() {
	for path, script := range c.scripts {
		println(path)

		jsonText, err := json.MarshalIndent(script, "", "  ")
		if err != nil {
			log.Println("Failed to marshal script to JSON!")
			return
		}

		println(string(jsonText))
	}
}

func (c *Compiler) compileDir(baseDir, dir string) error {
	// iterate over all files in dir and subdirectories which have the .plsd extension
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err := c.compileDir(baseDir, path.Join(dir, entry.Name())); err != nil {
				return err
			}
		} else if path.Ext(entry.Name()) == ".plsd" {
			if err := c.compileFile(baseDir, path.Join(dir, entry.Name())); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Compiler) compileFile(baseDir, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	relPath := strings.TrimPrefix(strings.ReplaceAll("\\", "/", file), strings.ReplaceAll("\\", "/", baseDir))
	relPath = strings.TrimPrefix(relPath, "/")

	is := antlr.NewInputStream(string(data))
	lexer := parser.NewProtoLSDLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewProtoLSDParser(stream)

	a := &analyzer{
		script: &script{
			Path:        relPath,
			Imports:     []*importStatement{},
			TypeAliases: map[string]string{},
			Enums:       map[string]*enum{},
			Messages:    map[string]*message{},
			Services:    map[string]*service{},
		},
		inGlobalScope: true,
	}

	antlr.ParseTreeWalkerDefault.Walk(a, p.ProtoLSD())

	c.scripts[relPath] = a.script

	return nil
}
