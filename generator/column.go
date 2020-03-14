package generator

import (
	"fmt"
)

/** TODO: Abstract class column */

type Column struct {
}

type Str struct {
	Column
}

// type Int struct {
// 	Column
// }

type InterfaceColumn interface {
	GenerateValue(InterfaceColumn) error
	// NextValue()
}

func (s Str) GenerateValue(i InterfaceColumn) error {
	fmt.Println("String")
	return nil
}
