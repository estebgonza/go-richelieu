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
	value string
	DefaultColumn
}

type IntColumn struct {
	value uint8
	DefaultColumn
}

type DateColumn struct {
	value time.Time
	DefaultColumn
}

type InterfaceColumn interface {
	GenerateValue(InterfaceColumn) (InterfaceColumn, error)
}

func (s StringColumn) GenerateValue(i InterfaceColumn) (InterfaceColumn, error) {
	fmt.Println(gofakeit.Word())
	s.value = gofakeit.Word()
	return s, nil
}

func (s IntColumn) GenerateValue(i InterfaceColumn) (InterfaceColumn, error) {
	fmt.Println(gofakeit.Uint8())
	s.value = gofakeit.Uint8()
	return s, nil
}

func (s DateColumn) GenerateValue(i InterfaceColumn) (InterfaceColumn, error) {
	fmt.Println(gofakeit.Date())
	s.value = gofakeit.Date()
	return s, nil
}
