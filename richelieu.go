package main

// TODO: Handle related cardinality (eg a column of description related to a column code)
// TODO: Handle unbalanced cardinality (eg a value represented on 80% of lines  )

// TODO: Extract tables row count from diag
// TODO: Parse showMemory.csv to estimate relatives cardinalities between indexed columns

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/estebgonza/go-richelieu/constants"
	"github.com/estebgonza/go-richelieu/creator"
	"github.com/estebgonza/go-richelieu/generator"
	"github.com/urfave/cli/v2"
)

const helpTemplate = `
Usage: {{.HelpName}} [command]

{{if .Commands}}Commands:

{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}
`

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(helpTemplate)
	app := cli.NewApp()
	app.Name = constants.AppName
	app.Usage = constants.AppDescription
	app.Version = constants.AppVersion

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Usage:   "Generate the dataset from the plan.json input",
			Aliases: []string{"g"},
			Action:  func(c *cli.Context) error { return generator.Generate() },
		},
		{
			Name:    "readFromSchema",
			Usage:   "Create a plan.json from db schema in createDataspace.txt",
			Aliases: []string{"rs"},
			Action:  func(c *cli.Context) error { return creator.CreateFromSchema() },
		},
		{
			Name:    "readFromColumn",
			Usage:   "Create a plan.json from column list argument",
			Aliases: []string{"rc"},
			Action:  func(c *cli.Context) error { return creator.CreateFromColumn(c.Args()) },
		},
		{
			Name:    "readFromDico",
			Usage:   "Update a plan.json from cardinality read in dictionary",
			Aliases: []string{"rd"},
			Action:  func(c *cli.Context) error { return creator.ReadCardinalityFromDictionaries() },
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("Command not found: %v\n", command)
		cli.ShowAppHelp(c)
	}

	log.Println("Starting...")
	startTime := time.Now()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done in", time.Since(startTime))
}
