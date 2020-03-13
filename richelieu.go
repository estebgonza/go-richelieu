package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	appName        string = "Richelieu"
	appDescription string = "Data generator that respects cardinality and schema structures#."
	appVersion     string = "0.1"
)

const helpTemplate = `
Usage: {{.HelpName}} [command]

{{if .Commands}}Commands:

{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}
`

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(helpTemplate)
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDescription
	app.Version = appVersion

	app.Commands = []*cli.Command{
		{
			Name:   "generate",
			Usage:  "Execute the generation plan",
			Action: generate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func generate(c *cli.Context) error {
	/*
	 Run plan in generator and returns
	*/
	return nil
}
