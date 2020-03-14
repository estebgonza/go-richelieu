package generator

import (
	"errors"
	"fmt"
)

type Rows struct {
	Rows []Row
}

type Row struct {
	Columns []Column
}

type Column struct {
	Type        string `json:"type"`
	Cardinality int    `json:"cardinality"`
}

type Plan struct {
	Rows    int      `json:"rows"`
	Columns []Column `json:"columns"`
}

func Execute(p *Plan) error {
	if err := validate(p); err != nil {
		return err
	}
	fmt.Println(p, p.Rows)
	// var row Row
	// var rows Rows
	// fmt.Println(rows)
	for i := 0; i < p.Rows; i++ {
		// if i%1024 == 0 && i != 0 {
		// 	flushRows(rows)
		// }
		for _, column := range p.Columns {
			// row = Row{}
			// fmt.Println(row, column)
			col := setType(column.Type)
			newCol, err := col.GenerateValue(col)
			fmt.Println(newCol)
			if err != nil {
				return err
			}
			//row.append(newCol)
		}
		// rows.appendRow(row)
	}
	// flushRows(rows)
	return nil
}

// func (r Row) appendCol(col Column)

// func (r Rows) appendRow(row Row) {
// 	r.append(row)
// }

func (r Rows) outputRows() {
	fmt.Println(r)
}

func flushRows(r Rows) error {
	r.outputRows()
	r.Rows = r.Rows[:0]
	r.outputRows()
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

func setType(t string) InterfaceColumn {

	switch t {
	case "STRING":
		return StringColumn{}
	case "INT":
		return IntColumn{}
	case "DATE":
		return DateColumn{}
	default:
		return DefaultColumn{}
	}
}
