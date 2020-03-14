package generator

import (
	"fmt"
)

/** TODO: Abstract class column
In default_column.go for now
*/

type Str struct {
	DefaultColumn
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
