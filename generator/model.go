package generator

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var ValidTypes []string = []string{"INT", "STRING", "FLOAT", "DATE"}

type Plan struct {
	Schemas []Schema `json:"schemas"`
}

type Schema struct {
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

type Table struct {
	Name    string   `json:"name"`
	Rows    int      `json:"rows"`
	Files   int      `json:"files"`
	Mode    string   `json:"mode,omitempty"` // Can be one of RANDOM, ALTERNATE, BLOCK
	Columns []Column `json:"columns"`
}

type Column struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Distinct    int      `json:"distinct"`
	Mode        string   `json:"mode,omitempty"` // Mode can be column specific
	Offset      int      `json:"offset,omitempty"`
	Start       string   `json:"start,omitempty"`
	End         string   `json:"end,omitempty"`
	Prefix      string   `json:"prefix,omitempty"`
	ValuesList  string   `json:"values,omitempty"`
	ValuesSlice []string `json:"-"` // Overload to store valuesList sliced
	FloatStart  float64  `json:"-"` // Pre-calculate float start
	DateStart   int64    `json:"-"` // Pre-calculate date start
	FloatStep   float64  `json:"-"` // Pre-calculate float step
	DateStep    int64    `json:"-"` // Pre-calculate date step
}

// Check that input type is supported
func ChecksSupportedType(t string) bool {
	for _, vt := range ValidTypes {
		if t == vt {
			return true
		}
	}
	return false
}

// Merge parts of plan
func MergePlanParts(plan *Plan, added *Plan) {
	if added == nil {
		return
	}
	for _, s := range added.Schemas {
		var i = plan.HasSchema(s.Name)
		if i < 0 { // Schema is not already in this plan, just add it
			plan.Schemas = append(plan.Schemas, s)
		} else { // Schema is already present, merge the tables
			mergeSchemaParts(&plan.Schemas[i], &s)
		}
	}

}

// Merge parts of schema
func mergeSchemaParts(schema *Schema, added *Schema) {
	if added == nil {
		return
	}
	for _, t := range added.Tables {
		var i = schema.HasTable(t.Name)
		if i < 0 { // Table is not already in this plan, just add it
			schema.Tables = append(schema.Tables, t)
		} else { // Table is already present, merge the tables
			mergeTableParts(&schema.Tables[i], &t)
		}
	}

}

// Merge parts of table
func mergeTableParts(table *Table, added *Table) {
	if added == nil {
		return
	}
	if added.Rows != 0 {
		table.Rows = added.Rows
	}
	if added.Files != 0 {
		table.Files = added.Files
	}
	if added.Mode != "" {
		table.Mode = added.Mode
	}
	if len(added.Columns) > 0 {
		table.Columns = added.Columns
	}
}

// Merge parts of column
func mergeColumnParts(column *Column, added *Column) {
	if added == nil {
		return
	}
	if added.Type != "" {
		column.Type = added.Type
	}
	if added.Offset != 0 {
		column.Offset = added.Offset
	}
	if added.Distinct != 0 {
		column.Distinct = added.Distinct
	}
	if added.Mode != "" {
		column.Mode = added.Mode
	}
	if added.Start != "" {
		column.Start = added.Start
	}
	if added.End != "" {
		column.End = added.End
	}
	if added.Prefix != "" {
		column.Prefix = added.Prefix
	}
	if added.ValuesList != "" {
		column.ValuesList = added.ValuesList
	}
}

func (plan *Plan) HasSchema(name string) int {
	for index := range plan.Schemas {
		if plan.Schemas[index].Name == name {
			return index
		}
	}
	return -1
}

func (schema *Schema) HasTable(name string) int {
	for index := range schema.Tables {
		if schema.Tables[index].Name == name {
			return index
		}
	}
	return -1
}

func (table *Table) HasColumn(name string) int {
	for index := range table.Columns {
		if table.Columns[index].Name == name {
			return index
		}
	}
	return -1
}

func (c *Column) getValue(line int, total int) string {
	var v int
	if c.Mode == "BLOCK" {
		v = int(float64(line) / (float64(total) / float64(c.Distinct)))
	} else if c.Mode == "RANDOM" {
		v = rand.Intn(c.Distinct)
	} else {
		v = line % c.Distinct
	}

	// Handle forced value list
	if v < len(c.ValuesSlice) {
		return c.ValuesSlice[v]
	}

	if c.Type == "FLOAT" {
		vv := c.FloatStart + c.FloatStep*float64(v)
		return strconv.FormatFloat(vv, 'f', -1, 32)
	} else if c.Type == "DATE" {
		vv := c.DateStart + c.DateStep*int64(v)
		tm := time.Unix(vv, 0)
		return tm.Local().String()
	} else { // INT, STRING
		return c.Prefix + strconv.Itoa(v+c.Offset)
	}
}

func ReadFromFile(planfile string) (*Plan, error) {
	var p Plan

	if _, err := os.Stat(planfile); os.IsNotExist(err) {
		return nil, nil // planfile doesnt exist
	}

	planFile, err := os.Open(planfile)
	if err != nil {
		return nil, errors.New("No plan.json found")
	}

	byteValue, _ := ioutil.ReadAll(planFile)
	json.Unmarshal(byteValue, &p)
	planFile.Close()

	// Control input plan
	if err := validate(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func WriteToFile(plan Plan, planfile string) error {

	// If planfile exist, we merge new plan with existing plan
	existingPlan, err := ReadFromFile(planfile)
	if err != nil {
		return err
	}
	if existingPlan != nil {
		MergePlanParts(&plan, existingPlan)
	}

	json, err := json.MarshalIndent(plan, "", "    ")
	if err != nil {
		return err
	}

	// Cosmetic corrections to have columns on one line (usefull for large schema)
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"type\"", "\"type\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"name\"", "\"name\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"distinct\"", "\"distinct\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"start\"", "\"start\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"end\"", "\"end\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                            \"values\"", "\"values\""))
	json = []byte(strings.ReplaceAll(string(json), "\n                        }", "}"))

	ioutil.WriteFile(planfile, json, 0644)

	return nil
}
