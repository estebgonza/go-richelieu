package generator

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// PlanColumn The plan per column as described in the input json file
type PlanColumn struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Distinct int    `json:"distinct"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

// Plan Plan is the plan as described in the input json file
type Plan struct {
	Rows        int          `json:"rows"`
	Files       int          `json:"files"`
	PlanColumns []PlanColumn `json:"columns"`
	Columns     []*Column    `json:"-"`
}

const (
	intType    = "INT"
	idIntType  = "ID_INT"
	floatType  = "FLOAT"
	dateType   = "DATE"
	stringType = "STRING"
)

// Execute Entrypoint of the generation plan
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
	wg := sync.WaitGroup{}
	for i := 0; i < p.Files; i++ {
		wg.Add(1)
		go func(i int) {
			fileName := "output/output_" + strconv.Itoa(i) + ".csv"
			csvFile, err := os.Create(fileName)
			if err != nil {
				log.Println(err)
			}
			csvWriter := csv.NewWriter(csvFile)
			for j := 0; j < p.Rows/p.Files; j++ {
				var row []string
				// Build the row
				for _, column := range p.Columns {
					// TODO use a master thread for cardinality management that listen to all the other threads and
					// change the c.currentValue accordingly
					column.nextValue()
					row = append(row, column.valueGenerator.getCurrentValue())
				}
				csvWriter.Write(row)
				if j%10000 == 0 && j != 0 {
					csvWriter.Flush()
				}
			}
			csvWriter.Flush()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nil
}

func initializeColumns(p *Plan) error {
	for _, planColumn := range p.PlanColumns {
		value, err := createValueGenerator(planColumn.Type)
		// TODO: Add a step calculator
		value.init(planColumn.Start)
		if err != nil {
			return err
		}
		rotBase := p.Rows / planColumn.Distinct
		rotMod := p.Rows % planColumn.Distinct
		name := planColumn.Name
		column := Column{valueGenerator: value, colName: name, rotationBase: rotBase, rotationMod: rotMod, count: rotBase, totCount: 0}
		p.Columns = append(p.Columns, &column)
	}
	return nil
}

// Validate Plan inputs.
// If validation fail returns an error.
func validate(p *Plan) error {
	rows := p.Rows
	if rows < 0 {
		return errors.New("Expected rows can't be negative")
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

// ChecksSupportedType Check that input type is supported by Richelieu by creating a temp instance of a Value
func ChecksSupportedType(t string) error {
	_, err := createValueGenerator(t)
	return err
}

func createValueGenerator(t string) (value, error) {
	var v value
	switch t {
	case intType:
		v = &intValue{}
	case idIntType:
		v = &idIntValue{}
	case floatType:
		v = &floatValue{}
	case dateType:
		v = &dateValue{}
	case stringType:
		v = &stringValue{}
	default:
		return nil, errors.New("Unsupported type " + t)
	}
	return v, nil
}
