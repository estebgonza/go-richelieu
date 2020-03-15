package generator

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v4"
)

type StringValue struct{}
type IntValue struct{}
type DateValue struct{}

type Value interface {
	GenerateValue() (string, error)
}

func (s StringValue) GenerateValue() (string, error) {
	return gofakeit.Word(), nil
}

func (s IntValue) GenerateValue() (string, error) {
	return strconv.Itoa(int(gofakeit.Uint8())), nil
}

func (s DateValue) GenerateValue() (string, error) {
	return gofakeit.Date().String(), nil
}
