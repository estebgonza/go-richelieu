package generator

import (
	"errors"
	"fmt"
)

type Plan struct {
	Rows    int `json:"rows"`
	Columns []struct {
		Type        string `json:"type"`
		Cardinality int    `json:"cardinality"`
	} `json:"columns"`
}

func Execute(p *Plan) error {
	if err := validate(p); err != nil {
		return err
	}
	// TODO: Execute the plan generation.
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
	for index, column := range p.Columns {
		cardinality := column.Cardinality
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
