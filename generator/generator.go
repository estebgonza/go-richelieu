package generator

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type PlanColumn struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Distinct int    `json:"distinct"`
}

type Plan struct {
	Rows        int          `json:"rows"`
	PlanColumns []PlanColumn `json:"columns"`
	Columns     []*Column    `json:"-"`
}

func Execute(p *Plan) error {
	if err := validate(p); err != nil {
		return err
	}
	// Initialize Column from PlanColumns
	initializeColumns(p)
	// Generate rows
	generate(p)

	return nil
}

func generate(p *Plan) error {
	csvFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	csvwriter := csv.NewWriter(csvFile)
	for i := 0; i < p.Rows; i++ {
		var row []string
		// Build the row
		for _, column := range p.Columns {
			row = append(row, column.nextValue())
		}
		csvwriter.Write(row)
	}
	csvwriter.Flush()
	return nil
}

func initializeColumns(p *Plan) {
	for _, planColumn := range p.PlanColumns {
		value := createValueGenerator(planColumn.Type)
		rot_base := p.Rows / planColumn.Distinct
		rot_mod := p.Rows % planColumn.Distinct
		name := planColumn.Name
		column := Column{valueGenerator: value, colName: name, rotationBase: rot_base, rotationMod: rot_mod, count: rot_base, totCount: 0}
		p.Columns = append(p.Columns, &column)
	}
}

// Validate Plan inputs.
// If validation fail returns an error.
func validate(p *Plan) error {
	rows := p.Rows
	if rows < 0 {
		return errors.New("Expected rows can't be negative.")
	}
	// Checks cardinalities for each columns
	for index, planColumn := range p.PlanColumns {
		cardinality := planColumn.Distinct
		if cardinality < 1 {
			m := fmt.Sprintf("Error. Column %d: cardinality can't be lower than 1.", index)
			return errors.New(m)
		}
		if cardinality > rows {
			m := fmt.Sprintf("Error. Column %d: cardinality can't be higher than number of rows (%d).", index, rows)
			return errors.New(m)
		}
	}
	return nil
}

func createValueGenerator(t string) Value {
	switch t {
	case "INT":
		return IntValue{}
	case "DATE":
		return DateValue{}
	case "STRING":
		return StringValue{}
	default:
		return StringValue{}
	}
}
