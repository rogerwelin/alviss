package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}

func validateRun(c *cli.Context) error {

	apiTmpl := &tmplData{
		APIProjectName:     c.String("project-name"),
		APIProtocol:        c.String("api-type"),
		APIEndpoints:       c.String("api-endpoint"),
		LambdaFunctionName: "helloworld",
		Language:           c.String("language"),
	}

	err := apiTmpl.bootstrapAPI()
	if err != nil {
		return err
	}

	return nil
}

func runCLI(args []string) {
	app := cli.NewApp()
	app.Name = "alviss"
	app.HelpName = "alviss"
	app.UsageText = "alviss [command] [command options] [arguments...]"
	app.EnableBashCompletion = true
	app.Usage = ""
	app.Version = "0.1.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Roger Welin",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "new-api",
			Usage: "Generates a new api project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "p, project-name",
					Usage:    "name of your API project",
					Required: true,
				},
				cli.GenericFlag{
					Name:  "t, api-type",
					Usage: "api type (only rest supported for now)",
					Value: &EnumValue{
						Enum:    []string{"rest"},
						Default: "rest",
					},
				},
				cli.GenericFlag{
					Name:  "e, api-endpoint",
					Usage: "which endpoint type (either regional, edge or private)",
					Value: &EnumValue{
						Enum:    []string{"regional", "edge", "private"},
						Default: "regional",
					},
				},
				cli.GenericFlag{
					Name:  "l, language",
					Usage: "which language for lambda to be used (go, node, python, ruby)",
					Value: &EnumValue{
						Enum:    []string{"go", "node", "python", "ruby"},
						Default: "node",
					},
				},
			},
			Action: validateRun,
		},
	}

	err := app.Run(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
