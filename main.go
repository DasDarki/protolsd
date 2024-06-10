package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"protolsd/compiler"
	"protolsd/util"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "protolsd",
		Usage:  "A protobuf superset on LSD. This is the CLI tool.",
		Action: compile,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "The directory to compile",
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"V"},
				Usage:       "Enable verbose output",
				Required:    false,
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"D"},
				Usage:       "Enable debug output",
				Required:    false,
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "quiet",
				Aliases:     []string{"Q"},
				Usage:       "Disable all output. This will overwrite verbose and debug flags.",
				Required:    false,
				DefaultText: "false",
			},
		},
		Commands: []*cli.Command{
			initCmd(),
			lspCompileCmd(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initCmd() *cli.Command {
	return &cli.Command{
		Name: "init",
		Action: func(c *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				log.Println("CRITICAL: Failed to get current working directory!")
				return err
			}

			if err := compiler.GenerateMinimalConfig(cwd); err != nil {
				log.Println("CRITICAL: Failed to generate minimal config!")
				return err
			}

			return nil
		},
	}
}

func lspCompileCmd() *cli.Command {
	return &cli.Command{
		Name: "lsp-compile",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "workspace",
				Aliases: []string{"w"},
				Usage:   "The workspace to compile",
			},
		},
		Action: lspCompile,
	}
}

func compile(c *cli.Context) error {
	logger := util.NewLogger(c.Bool("quiet"), c.Bool("verbose"), c.Bool("debug"))

	dir := ""
	if c.String("dir") != "" {
		dir = c.String("dir")
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			logger.Crit("Failed to get current working directory!")
			return err
		}

		dir = cwd
	}

	config, err := compiler.LoadConfig(dir)
	if err != nil {
		logger.Crit("Failed to load config! Do you have a protolsd.toml file?")
		return err
	}

	compiler := compiler.NewCompiler(config, dir, false, logger)
	if err := compiler.Compile(); err != nil {
		logger.Crit("Failed to compile!")

		return err
	}

	log.Println("INFO: Done!")

	return nil
}

func lspCompile(c *cli.Context) error {
	workspace := c.String("workspace")

	var config *compiler.Config
	confFile := path.Join(workspace, "protolsd.toml")

	if _, err := os.Stat(confFile); err == nil {
		var err error
		config, err = compiler.LoadConfig(workspace)
		if err != nil {
			log.Fatalf("Failed to load config! Do you have a protolsd.toml file? %v", err)
		}
	} else {
		config = &compiler.Config{
			InputDir:     "src",
			EnvFile:      util.Ptr(""),
			OutputDir:    util.Ptr(""),
			OutputType:   util.Ptr(compiler.OutputTypeLSD),
			OrderPersist: util.Ptr(true),
			WithGRPC:     util.Ptr(false),
			Protobuf: &compiler.ConfigProtobufSection{
				AutoDL:        util.Ptr(false),
				Version:       util.Ptr("latest"),
				Package:       util.Ptr("main"),
				Targets:       []string{},
				TargetOptions: map[string]map[string]string{},
				Options:       map[string]string{},
			},
		}
	}

	comp := compiler.NewCompiler(config, workspace, true, util.NewLogger(true, false, false))
	compReuslt := map[string]any{}

	if err := comp.Compile(); err != nil {
		compReuslt["success"] = false
		compReuslt["error"] = err
	} else {
		compReuslt["success"] = true
		compReuslt["result"] = comp.GetCompiledPackage()
	}

	jsonResult, err := json.Marshal(compReuslt)
	if err != nil {
		log.Fatalf("Failed to marshal result! %v", err)
	}

	println(string(jsonResult))

	return nil
}
