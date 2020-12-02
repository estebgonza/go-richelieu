package generator

import "strconv"

type intValue struct {
	currentValue int64
	step         int64
}

func (v *intValue) generateValue() {
	v.currentValue += v.step
}

func (v *intValue) init(i string) {
	v.currentValue, _ = strconv.ParseInt(i, 10, 64)
	v.step = 1
}

func (v intValue) getCurrentValue() string {
	return strconv.FormatInt(v.currentValue, 10)
}
