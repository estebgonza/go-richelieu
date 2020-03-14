package generator

import (
	"fmt"
)

type DefaultColumn struct {
}

func (d DefaultColumn) GenerateValue(i InterfaceColumn) (string, error) {
	fmt.Println("DEFAULT STRING")
	return "DEFAULT STRING", nil
}
