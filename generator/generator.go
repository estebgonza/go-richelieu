package generator

import "fmt"

type Plan struct {
	Rows    int `json:"rows"`
	Columns []struct {
		Type        string `json:"type"`
		Cardinality int    `json:"cardinality"`
	} `json:"columns"`
}

func Execute(p *Plan) error {
	/**
	TODO:
	Check if input Plan is valid.
	Execute the plan generation.
	*/
	fmt.Println("Executed")
	return nil
}
