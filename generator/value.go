package generator

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v4"
)

type StringValue struct{}
type IntValue struct{}
type DateValue struct{}

type Value interface {
	GenerateValue(b string, c uint64) (string, error)
}

func (s StringValue) GenerateValue(b string, c uint64) (string, error) {
	return b + "_" + strconv.Itoa(int(c)), nil
}

func (s IntValue) GenerateValue(b string, c uint64) (string, error) {
	return strconv.Itoa(int(c)), nil
}

func (s DateValue) GenerateValue(b string, c uint64) (string, error) {
	var date time.Time
	var ret bytes.Buffer
	date = gofakeit.Date()
	_, err := fmt.Fprintf(&ret, "%v-%02d-%02d %02d:%02d:%02d", date.Year(), int(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second())
	if err != nil {
		return "", err
	}
	return ret.String(), nil
}
