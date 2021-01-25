package generator

import "strconv"

type stringValue struct {
	prefix       string
	currentValue string
	step         int64
	currentStep  int64
}

func (v *stringValue) generateValue() {
	v.currentStep = v.currentStep + v.step
	v.currentValue = v.prefix + "_" + strconv.FormatInt(v.currentStep, 10)
}

func (v *stringValue) init(i string) {
	v.prefix = i
	v.currentStep = 1
	v.currentValue = v.prefix + "_" + strconv.FormatInt(v.currentStep, 10)
	v.step = 1
}

func (v stringValue) getCurrentValue() string {
	return v.currentValue
}
