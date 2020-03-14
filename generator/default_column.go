package generator

import (
	"fmt"
)

type DefaultColumn struct {
	value string
}

func (d DefaultColumn) GenerateValue(i InterfaceColumn) (InterfaceColumn, error) {
	fmt.Println("DEFAULT STRING")
	d.value = "DEFAULT STRING"
	return d, nil
}
