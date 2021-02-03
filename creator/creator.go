package creator

import (
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/estebgonza/go-richelieu/constants"
	"github.com/estebgonza/go-richelieu/generator"
	"github.com/urfave/cli/v2"
)

// Note: Parsing a schema file manually with regexp
// A proper sql parser would be much better, but couldnt find one that handle hive
// "vitess.io/vitess/go/vt/sqlparser" works very well but doesnt parse many elements of hive

// Read a text file containing create table queries and initiate a plan.json with it
func CreateFromSchema() error {

	// Read schema input sql file
	content, err := ioutil.ReadFile(constants.DefaultSchemaFile)
	if err != nil {
		return err
	}
	text := string(content)

	// Check create table count
	cntTable := strings.Count(text, "CREATE DIMENSION TABLE") + strings.Count(text, "CREATE TABLE")
	log.Println("Found " + strconv.Itoa(cntTable) + " tables to extract")

	text = cleansingSchemaFile(text)
	allQueries := strings.Split(text, ";")
	// Remove empty last part
	if !strings.Contains(allQueries[len(allQueries)-1], "(") {
		allQueries = allQueries[:len(allQueries)-1]
	}

	// Check that number of semicolon now match table creation
	if len(allQueries) != cntTable {
		return errors.New("Error while cleansing schema file. Incorrect table count")
	}

	// Loop through queries to extract tables creation
	plan := generator.Plan{}
	for _, q := range allQueries {
		q = strings.ReplaceAll(q, "\n", "")
		// Search parenthesis (start of field list)
		startField := strings.Index(q, "(")
		if startField == -1 {
			return errors.New("Error while cleansing schema file. Cant find table fields")
		}
		name := strings.ReplaceAll(q[:startField], " ", "")
		endField := strings.LastIndex(q, ")")
		q = q[startField+1 : endField]
		q = strings.ReplaceAll(q, ",,", ",")

		// Split, removing empty parts (due to cleansing of main index)
		splitFn := func(c rune) bool { return c == ',' }
		allFields := strings.FieldsFunc(q, splitFn)

		names := strings.Split(name, ".")

		// Add the default schema for table withotu schema
		if len(names) == 1 {
			names = strings.Split("default."+name, ".")
		}
		// Check parsing is correct (schema.table)
		if len(names) != 2 {
			return errors.New("Error reading schema and table name " + names[1])
		}

		table := generator.Table{Name: names[1]}
		for _, c := range allFields {
			v := strings.Split(c, " ")
			if len(v) != 2 {
				return errors.New("Error reading field from table " + names[1])
			}
			table.Columns = append(table.Columns, generator.Column{Name: v[0], Type: v[1], Distinct: 1})
		}
		schema := generator.Schema{Name: names[0], Tables: []generator.Table{table}}
		newPlan := generator.Plan{Schemas: []generator.Schema{schema}}
		generator.MergePlanParts(&plan, &newPlan)
	}

	updatePlan(&plan)

	return generator.WriteToFile(plan, constants.DefaultPlanFile)
}

// Fill files and rows if null
func updatePlan(p *generator.Plan) {
	for s := range p.Schemas { // For each schema
		for t := range p.Schemas[s].Tables { // For each table
			if p.Schemas[s].Tables[t].Rows == 0 {
				p.Schemas[s].Tables[t].Rows = 100
			}
			if p.Schemas[s].Tables[t].Files == 0 {
				p.Schemas[s].Tables[t].Files = 1
			}
		}
	}
}

// Custom parser to cleanse a database create schema file
func cleansingSchemaFile(text string) string {

	// Remove create schema
	r, _ := regexp.Compile("CREATE SCHEMA (.*);")
	text = r.ReplaceAllString(text, "")

	// Simplify create table statement
	text = strings.ReplaceAll(text, "CREATE DIMENSION TABLE ", "CREATE TABLE ")
	text = strings.ReplaceAll(text, "CREATE TABLE IF NOT EXISTS ", "CREATE TABLE ")

	// Cleanse field prefixed with --
	text = strings.ReplaceAll(text, "-- ", "")

	// Remove table source (eg CREATE TABLE table1 FROM srcTable IN 'jdbc:hive2://table')
	r, _ = regexp.Compile("FROM (.*) IN '(.*)'")
	text = r.ReplaceAllString(text, "")

	// Simplify numeric types
	r, _ = regexp.Compile("DOUBLE\\(([0-9]+),([0-9]+)\\)")
	text = r.ReplaceAllString(text, "DOUBLE")

	// Remove ALTER TABLE
	r, _ = regexp.Compile("ALTER TABLE (.*);")
	text = r.ReplaceAllString(text, "")

	// Remove MAIN INDEX
	r, _ = regexp.Compile("MAIN INDEX (.*)\\((.*)\\)")
	text = r.ReplaceAllString(text, "")

	// Remove Table PROPERTIES
	r, _ = regexp.Compile("PROPERTIES \\((.*)\\)")
	text = r.ReplaceAllString(text, "")

	// Replace AS round with DOUBLE
	r, _ = regexp.Compile("AS round\\((.*)\\)")
	text = r.ReplaceAllString(text, "DOUBLE")

	// Replace AS if with DOUBLE
	r, _ = regexp.Compile("AS if\\((.*)\\)")
	text = r.ReplaceAllString(text, "DOUBLE")

	// Remove CREATE TABLE
	text = strings.ReplaceAll(text, "CREATE TABLE ", "")

	return text
}

// Read a list of types from arguments (eg: INT, STRING, INT) and initiate a plan.json with it
func CreateFromColumn(args cli.Args) error {

	if args.Len() == 0 {
		return errors.New("Please specify columns type to init a generation plan")
	}

	typeList := strings.ReplaceAll(args.Get(0), " ", "")
	cols := strings.Split(typeList, ",")
	var columns []generator.Column
	for index, t := range cols {
		if !generator.ChecksSupportedType(t) {
			return errors.New("Unsupported column type " + t)
		}
		var pc generator.Column
		pc.Name = strings.ToLower(strconv.Itoa(index) + "_" + t)
		pc.Distinct = 1
		pc.Type = t
		columns = append(columns, pc)
	}

	var table = generator.Table{Name: "table1", Rows: 10000, Files: 1, Columns: columns}
	var schema = generator.Schema{Name: "schema1", Tables: []generator.Table{table}}
	var plan = generator.Plan{Schemas: []generator.Schema{schema}}

	return generator.WriteToFile(plan, constants.DefaultPlanFile)
}

// Read a dictionary to extract columns cardinality
func ReadCardinalityFromDictionaries() error {

	// Load current plan
	p, err := generator.ReadFromFile(constants.DefaultPlanFile)
	if err != nil {
		return err
	}

	// Read dictionary
	csvFile, err := os.Open(constants.DefaultDictionaryFile)
	if err != nil {
		return err
	}
	csvReader := csv.NewReader(csvFile)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		schemaIndex := p.HasSchema(record[1])
		if schemaIndex < 0 {
			continue
		}
		tableIndex := p.Schemas[schemaIndex].HasTable(record[2])
		if tableIndex < 0 {
			continue
		}
		columnIndex := p.Schemas[schemaIndex].Tables[tableIndex].HasColumn(record[3])
		if columnIndex < 0 {
			continue
		}
		card, err := strconv.Atoi(record[4])
		if err != nil {
			continue
		}
		if card < 1 {
			continue
		}
		p.Schemas[schemaIndex].Tables[tableIndex].Columns[columnIndex].Distinct = card
	}

	// Delete old plan, to prevent a merge
	_ = os.Remove(constants.DefaultPlanFile)

	// Rewrite plan
	return generator.WriteToFile(*p, constants.DefaultPlanFile)
}

func ReadTableCounts() error {

	// Load current plan
	p, err := generator.ReadFromFile(constants.DefaultPlanFile)
	if err != nil {
		return err
	}

	// Read dictionary
	csvFile, err := os.Open(constants.DefaultTableCountFile)
	if err != nil {
		return err
	}
	csvReader := csv.NewReader(csvFile)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		schemaIndex := p.HasSchema(record[0])
		if schemaIndex < 0 {
			continue
		}
		tableIndex := p.Schemas[schemaIndex].HasTable(record[1])
		if tableIndex < 0 {
			continue
		}
		rowCount, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}
		if rowCount < 1 {
			continue
		}
		p.Schemas[schemaIndex].Tables[tableIndex].Rows = rowCount
		if rowCount > 1000000 && p.Schemas[schemaIndex].Tables[tableIndex].Files == 1 {
			p.Schemas[schemaIndex].Tables[tableIndex].Files = 4
		}
	}

	// Delete old plan, to prevent a merge
	_ = os.Remove(constants.DefaultPlanFile)

	// Rewrite plan
	return generator.WriteToFile(*p, constants.DefaultPlanFile)
}
