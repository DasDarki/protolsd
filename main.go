package main

import (
	"log"
	"os"
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
			lspCmd(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	} else {
		log.Println("INFO: Done!")
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

	return nil
}

func lspCmd() *cli.Command {
	return &cli.Command{
		Name:        "lsp",
		Description: "Start a new language server protocol instance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dir",
				Aliases:  []string{"d"},
				Usage:    "The directory to compile",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

			return nil
		},
	}
}
