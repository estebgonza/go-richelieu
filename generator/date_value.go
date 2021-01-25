package generator

import (
	"log"
	"time"
)

type dateValue struct {
	currentValue time.Time
	step         time.Duration
}

func (v *dateValue) generateValue() {
	v.currentValue = v.currentValue.Add(v.step)
}

func (v *dateValue) init(i string) {
	var err error
	v.currentValue, err = time.Parse("2006-01-02 15:04:05", i)
	if err != nil {
		log.Fatalln(err)
	}
	v.step = 1000000000 // 1B nanoseconds = 1s
}

func (v dateValue) getCurrentValue() string {
	return v.currentValue.String()
}
