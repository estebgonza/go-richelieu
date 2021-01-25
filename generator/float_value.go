package generator

import "strconv"

type floatValue struct {
	currentValue float64
	step         float64
}

func (v *floatValue) generateValue() {
	v.currentValue += v.step
}

func (v *floatValue) init(i string) {
	v.currentValue, _ = strconv.ParseFloat(i, 64)
	v.step = 0.1
}

func (v floatValue) getCurrentValue() string {
	return strconv.FormatFloat(v.currentValue, 'f', 2, 64)
}
