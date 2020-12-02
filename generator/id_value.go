package generator

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v4"
)

type idIntValue struct {
	currentValue uint64
}

/*
 * All new types need all the methods in interface value in value.go
 */

func (v *idIntValue) generateValue() {
	v.currentValue = gofakeit.Uint64()
}

func (v *idIntValue) init(i string) {
	v.currentValue = gofakeit.Uint64()
}

func (v idIntValue) getCurrentValue() string {
	return strconv.FormatUint(v.currentValue, 10)
}
