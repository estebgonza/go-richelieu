package generator

import (
	"fmt"
)

type DefaultColumn struct {
}

func (d DefaultColumn) GenerateValue(InterfaceColumn) error {
	fmt.Println("STRING")
	return nil
}
