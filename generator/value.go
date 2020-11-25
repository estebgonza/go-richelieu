package generator

import (
	"log"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v4"
)

type stringValue struct {
	prefix       string
	currentValue string
	step         int64
	currentStep  int64
}

type intValue struct {
	currentValue int64
	step         int64
}

type dateValue struct {
	currentValue time.Time
	step         time.Duration
}

type idIntValue struct {
	currentValue uint64
}

type value interface {
	getCurrentValue() string
	generateValue()
	init(i string)
}

func (v *idIntValue) generateValue() {
	v.currentValue = gofakeit.Uint64()
}

func (v *intValue) generateValue() {
	v.currentValue += v.step
}

func (v *stringValue) generateValue() {
	v.currentStep = v.currentStep + v.step
	v.currentValue = v.prefix + "_" + strconv.FormatInt(v.currentStep, 10)
}

func (v *dateValue) generateValue() {
	v.currentValue.Add(v.step)
}

func (v *idIntValue) init(i string) {
	v.currentValue = gofakeit.Uint64()
}

func (v *intValue) init(i string) {
	v.currentValue, _ = strconv.ParseInt(i, 10, 64)
	v.step = 1
}

func (v *stringValue) init(i string) {
	v.prefix = i
	v.currentStep = 1
	v.currentValue = v.prefix + "_" + strconv.FormatInt(v.currentStep, 10)
	v.step = 1
}

func (v *dateValue) init(i string) {
	var err error
	v.currentValue, err = time.Parse("2006-01-02 15:04:05", i)
	if err != nil {
		log.Fatalln(err)
	}
	v.step = 1000000 // 1M nanoseconds = 1s
}

func (v idIntValue) getCurrentValue() string {
	return strconv.FormatUint(v.currentValue, 10)
}

func (v intValue) getCurrentValue() string {
	return strconv.FormatInt(v.currentValue, 10)
}

func (v stringValue) getCurrentValue() string {
	return v.currentValue
}

func (v dateValue) getCurrentValue() string {
	return v.currentValue.String()
}
