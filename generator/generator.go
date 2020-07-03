package generator

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type PlanColumn struct {
	Type        string `json:"type"`
	Cardinality int    `json:"cardinality"`
}

type Plan struct {
	Rows        int          `json:"rows"`
	PlanColumns []PlanColumn `json:"columns"`
	Columns     []*Column    `json:"-"`
}

const (
	INT_TYPE    = "INT"
	DATE_TYPE   = "DATE"
	STRING_TYPE = "STRING"
)

func Execute(p *Plan) error {
	if err := validate(p); err != nil {
		return err
	}
	// Initialize Column from PlanColumns
	if err := initializeColumns(p); err != nil {
		return err
	}
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

func initializeColumns(p *Plan) error {
	for _, planColumn := range p.PlanColumns {
		value, err := createValueGenerator(planColumn.Type)
		if err != nil {
			return err
		}
		rot_base := p.Rows / planColumn.Cardinality
		rot_mod := p.Rows % planColumn.Cardinality
		column := Column{valueGenerator: value, rotation_base: rot_base, rotation_mod: rot_mod, count: rot_base}
		p.Columns = append(p.Columns, &column)
	}
	return nil
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
		cardinality := planColumn.Cardinality
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

func ChecksSupportedType(t string) error {
	_, err := createValueGenerator(t)
	return err
}

func createValueGenerator(t string) (val Value, err error) {
	switch t {
	case INT_TYPE:
		return IntValue{}, nil
	case DATE_TYPE:
		return DateValue{}, nil
	case STRING_TYPE:
		return StringValue{}, nil
	default:
		return nil, errors.New("Unsupported type " + t)
	}
}
