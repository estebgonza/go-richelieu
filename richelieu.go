package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/estebgonza/go-richelieu/generator"
	"github.com/urfave/cli/v2"
)

const (
	appName         string = "Richelieu"
	appDescription  string = "Data generator that respects cardinality and schema structures#."
	appVersion      string = "0.1"
	defaultPlanFile string = "plan.json"
)

const helpTemplate = `
Usage: {{.HelpName}} [command]

{{if .Commands}}Commands:

{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}
`

func main() {
	log.SetFlags(0)
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
		{
<<<<<<< HEAD
			Name:   "create",
			Usage:  "Create a generation plan",
			Action: create,
=======
			Name:   "serve",
			Usage:  "Execute the generation plan",
			Action: serve,
>>>>>>> 47055457e1e99e1e1fe83bf5f9f0f432fac5ae80
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serve(c *cli.Context) error {
	err := errors.New("Not implemented")
	return err
}

func generate(c *cli.Context) error {
	var planFile *os.File
	var byteValue []byte
	var p generator.Plan
	gofakeit.Seed(time.Now().UnixNano())
	planFile, err := os.Open(defaultPlanFile)
	if err != nil {
		return errors.New("No plan.json found.")
	}
	byteValue, _ = ioutil.ReadAll(planFile)
	json.Unmarshal(byteValue, &p)
	errExec := generator.Execute(&p)
	if errExec != nil {
		return errExec
	}
	fmt.Printf("Done. %d rows just generated.\n", p.Rows)
	return nil
}

func create(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return errors.New("Please specify columns type to init a generation plan.")
	}
	cols := strings.Split(c.Args().Get(0), ",")
	for _, t := range cols {
		if err := generator.ChecksSupportedType(t); err != nil {
			return err
		}
	}
	return nil
}
