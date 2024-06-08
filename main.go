package main

import (
	"log"
	"os"
	"protolsd/compiler"

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
		},
		Commands: []*cli.Command{
			initCmd(),
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
	dir := ""
	if c.String("dir") != "" {
		dir = c.String("dir")
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			log.Println("CRITICAL: Failed to get current working directory!")
			return err
		}

		dir = cwd
	}

	config, err := compiler.LoadConfig(dir)
	if err != nil {
		log.Println("CRITICAL: Failed to load config! Do you have a protolsd.toml file?")
		return err
	}

	compiler := compiler.NewCompiler(config, dir)
	if err := compiler.Compile(); err != nil {
		log.Println("CRITICAL: Failed to compile!")
		return err
	}

	return nil
}
