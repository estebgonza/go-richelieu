package generator

import (
	"fmt"
)

type DefaultColumn struct {
}

func (d DefaultColumn) GenerateValue(i InterfaceColumn) error {
	fmt.Println("STRING")
	return nil
}