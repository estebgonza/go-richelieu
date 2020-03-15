package generator

import (
	"errors"
	"fmt"
)

type PlanColumn struct {
	Type        string `json:"type"`
	Cardinality int    `json:"cardinality"`
}

type Plan struct {
	Rows        int          `json:"rows"`
	PlanColumns []PlanColumn `json:"columns"`
	Columns     []Column     `json:"-"`
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

func generate(p *Plan) {
	for i := 0; i < p.Rows; i++ {
		for _, column := range p.Columns {
			fmt.Println(column.nextValue())
		}
	}
}

func initializeColumns(p *Plan) {
	for _, planColumn := range p.PlanColumns {
		value := createValueGenerator(planColumn.Type)
		column := Column{valueGenerator: value}
		p.Columns = append(p.Columns, column)
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
