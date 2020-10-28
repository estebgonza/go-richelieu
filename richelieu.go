package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
			Name:   "create",
			Usage:  "Create a generation plan",
			Action: create,
		},
	}
	log.Println("Starting generation...")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done generating")
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
	var columns []generator.PlanColumn
	for index, t := range cols {
		if err := generator.ChecksSupportedType(t); err != nil {
			return err
		}
		var pc generator.PlanColumn
		pc.Name = strings.ToLower(strconv.Itoa(index) + "_" + t)
		pc.Distinct = 5
		pc.Type = t
		columns = append(columns, pc)
	}

	var plan generator.Plan
	plan.Rows = 10000
	plan.PlanColumns = columns
	b, _ := json.Marshal(plan)
	ioutil.WriteFile(defaultPlanFile, b, 0644)
	return nil
}
