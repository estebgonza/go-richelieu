package generator

import (
	"reflect"
	"testing"
)

func TestFloatValueInit1(t *testing.T) {
	var r1 floatValue
	var r2 floatValue
	var e float64 = 1.00
	initValue := "1.00"
	r1.init(initValue)
	r2.init(initValue)
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(e) {
		t.Errorf("Expected type is %v. Got %v", reflect.TypeOf(e), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue != r2.currentValue {
		t.Errorf("r1 = %v. r2 = %v. Should be the same", r1.currentValue, r2.currentValue)
	}
	if r1.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r1.currentValue)
	}
	if r2.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r2.currentValue)
	}
}

func TestFloatValueInit2(t *testing.T) {
	var etype float64
	var r1 floatValue
	var r2 floatValue
	initValue1 := "1.00"
	initValue2 := "2.00"
	r1.init(initValue1)
	r2.init(initValue2)
	var e1 float64 = 1
	var e2 float64 = 2
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(etype) {
		t.Errorf("Expected type is %v. Got %v", reflect.TypeOf(etype), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue != e1 {
		t.Errorf("r1 KO. Expected %v. Got %v", r1.currentValue, e1)
	}
	if r2.currentValue != e2 {
		t.Errorf("r2 KO. Expected %v. Got %v", r2.currentValue, e2)
	}
	if r1.currentValue == r2.currentValue {
		t.Errorf("%v == %v. Should be different", r1.currentValue, r2.currentValue)
	}
}

func TestFloatValueGenerate1(t *testing.T) {
	var r floatValue
	initValue := "1"
	r.init(initValue)
	var e float64 = 1.10
	r.generateValue()
	if r.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r.currentValue)
	}
}

func TestFloatValueGetCurrentValue1(t *testing.T) {
	var r floatValue
	initValue := "1"
	r.init(initValue)
	var e string = "1.00"
	ret := r.getCurrentValue()
	if ret != e {
		t.Errorf("Expected %v. Got %v", e, ret)
	}
}
