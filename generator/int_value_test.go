package generator

import (
	"reflect"
	"testing"
)

func TestIntValueInit1(t *testing.T) {
	var r1 intValue
	var r2 intValue
	var e int64 = 1
	initValue := "1"
	r1.init(initValue)
	r2.init(initValue)
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(e) {
		t.Errorf("Expected type is %v. Got %v", reflect.TypeOf(e), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue != r2.currentValue {
		t.Errorf("r1 = %d. r2 = %d. Should be the same", r1.currentValue, r2.currentValue)
	}
	if r1.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r1.currentValue)
	}
	if r2.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r2.currentValue)
	}
}

func TestIntValueInit2(t *testing.T) {
	var etype int64
	var r1 intValue
	var r2 intValue
	initValue1 := "1"
	initValue2 := "2"
	r1.init(initValue1)
	r2.init(initValue2)
	var e1 int64 = 1
	var e2 int64 = 2
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
		t.Errorf("%d == %d. Should be different", r1.currentValue, r2.currentValue)
	}
}

func TestIntValueGenerate1(t *testing.T) {
	var r intValue
	initValue := "1"
	r.init(initValue)
	var e int64 = 2
	r.generateValue()
	if r.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r.currentValue)
	}
}

func TestIntValueGetCurrentValue1(t *testing.T) {
	var r intValue
	initValue := "1"
	r.init(initValue)
	var e string = "1"
	ret := r.getCurrentValue()
	if ret != e {
		t.Errorf("Expected %v. Got %v", e, ret)
	}
}
