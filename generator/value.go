package generator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v4"
)

type StringValue struct {
	prefix       string
	step         int64
	currentValue int64 // Current value of the suffix
}

type IntValue struct {
	currentValue int64
	step         int64
}

type DateValue struct {
	currentValue time.Time
	step         time.Time
}

type IdIntValue struct {
	currentValue uint64
}

type Value interface {
	GenerateValue() (string, error)
}

func (s IdIntValue) GenerateValue() (string, error) {
	return strconv.Itoa(int(gofakeit.Uint64())), nil
}

func (s StringValue) GenerateValue() (string, error) {
	return s.prefix + "_" + strconv.Itoa(int(s.currentValue)+int(s.step)), nil
}

func (s IntValue) GenerateValue() (string, error) {
	return strconv.FormatInt(s.currentValue+s.step, 10), nil
}

func (s DateValue) GenerateValue() (string, error) {
	var date time.Time
	date = gofakeit.Date()
	return fmt.Sprintf("%v-%02d-%02d %02d:%02d:%02d", date.Year(), int(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second()), nil
}
