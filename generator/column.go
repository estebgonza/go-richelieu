package generator

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v4"
)

/** TODO: Abstract class column
In default_column.go for now
*/

type StringColumn struct{}
type IntColumn struct{}
type DateColumn struct{}

type InterfaceColumn interface {
	GenerateValue(InterfaceColumn) (string, error)
}

func (s StringColumn) GenerateValue(i InterfaceColumn) (string, error) {
	return gofakeit.Word(), nil
}

func (s IntColumn) GenerateValue(i InterfaceColumn) (string, error) {
	return strconv.Itoa(int(gofakeit.Uint8())), nil
}

func (s DateColumn) GenerateValue(i InterfaceColumn) (string, error) {
	return gofakeit.Date().String(), nil
}
