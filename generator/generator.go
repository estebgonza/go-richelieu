package generator

type Plan struct {
	Rows    int `json:"rows"`
	Columns []struct {
		Type        string `json:"type"`
		Cardinality int    `json:"cardinality"`
	} `json:"columns"`
}
