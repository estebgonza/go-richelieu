package generator

import (
	"reflect"
	"testing"
)

func TestStringValueInit1(t *testing.T) {
	var r1 stringValue
	var r2 stringValue
	var e string = "test_1"
	initValue := "test"
	r1.init(initValue)
	r2.init(initValue)
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(e) {
		t.Errorf("Type is not %v. Is %v", reflect.TypeOf(e), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue != r2.currentValue {
		t.Errorf("r1 = %v. r2 = %v. Should be the same", r1.currentValue, r2.currentValue)
	}
	if r1.currentValue != e {
		t.Errorf("r1 KO. Expected %v. Got %v", r1.currentValue, e)
	}
	if r2.currentValue != e {
		t.Errorf("r2 KO. Expected %v. Got %v", r2.currentValue, e)
	}
}

func TestStringValueInit2(t *testing.T) {
	var etype string
	var r1 stringValue
	var r2 stringValue
	initValue1 := "test1"
	initValue2 := "test2"
	r1.init(initValue1)
	r2.init(initValue2)
	var e1 string = "test1_1"
	var e2 string = "test2_1"
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(etype) {
		t.Errorf("Type is not %v. Is %v", reflect.TypeOf(etype), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue != e1 {
		t.Errorf("r1 KO. Expected %v. Got %v", e1, r1.currentValue)
	}
	if r2.currentValue != e2 {
		t.Errorf("r2 KO. Expected %v. Got %v", e2, r2.currentValue)
	}
	if r1.currentValue == r2.currentValue {
		t.Errorf("%v == %v. Should be different", r1.currentValue, r2.currentValue)
	}
}

func TestStringValueGenerate1(t *testing.T) {
	var r stringValue
	initValue := "test"
	r.init(initValue)
	var e string = "test_2"
	r.generateValue()
	if r.currentValue != e {
		t.Errorf("Expected %v, got %v", e, r.currentValue)
	}
}

func TestStringValueGetCurrentValue1(t *testing.T) {
	var r stringValue
	initValue := "test"
	r.init(initValue)
	var e string = "test_1"
	ret := r.getCurrentValue()
	if ret != e {
		t.Errorf("Expected %v, got %v", e, ret)
	}
}
