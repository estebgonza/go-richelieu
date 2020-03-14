package generator

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v4"
	"time"
)

/** TODO: Abstract class column
In default_column.go for now
*/

type StringColumn struct {
	DefaultColumn
}

type IntColumn struct {
	DefaultColumn
}

type DateColumn struct {
	DefaultColumn
}

type InterfaceColumn interface {
	GenerateValue(InterfaceColumn) (string, error)
}

func (s StringColumn) GenerateValue(i InterfaceColumn) (string, error) {
	fmt.Println(gofakeit.Word())
	return gofakeit.Word(), nil
}

func (s IntColumn) GenerateValue(i InterfaceColumn) (uint8, error) {
	fmt.Println(gofakeit.Uint8())
	return gofakeit.Uint8(), nil
}

func (s DateColumn) GenerateValue(i InterfaceColumn) (time.Time, error) {
	fmt.Println(gofakeit.Date())
	return gofakeit.Date(), nil
}
