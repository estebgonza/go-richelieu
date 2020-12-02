package generator

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIdValueInit1(t *testing.T) {
	var etype uint64
	var r1 idIntValue
	var r2 idIntValue
	initValue := "1"
	r1.init(initValue)
	r2.init(initValue)
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(etype) {
		t.Errorf("Expected type is %v. Got %v", reflect.TypeOf(etype), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue == r2.currentValue {
		t.Errorf("%d == %d. Should be different", r1.currentValue, r2.currentValue)
	}
}

func TestIdValueInit2(t *testing.T) {
	var etype uint64
	var r1 idIntValue
	var r2 idIntValue
	initValue1 := "1"
	initValue2 := "2"
	r1.init(initValue1)
	r2.init(initValue2)
	if reflect.TypeOf(r1.currentValue) != reflect.TypeOf(etype) {
		t.Errorf("Expected type is %v. Got %v", reflect.TypeOf(etype), reflect.TypeOf(r1.currentValue))
	}
	if r1.currentValue == r2.currentValue {
		t.Errorf("%d == %d. Should be different", r1.currentValue, r2.currentValue)
	}
}

func TestIdValueGenerate1(t *testing.T) {
	var r idIntValue
	initValue := "1"
	r.init(initValue)
	ne := r.currentValue
	r.generateValue()
	if r.currentValue == ne {
		t.Errorf("Id is the same after a generate, should be different. Before: %v, After %v", ne, r.currentValue)
	}
}

func TestIdValueGetCurrentValue1(t *testing.T) {
	var r idIntValue
	initValue := "1"
	r.init(initValue)
	e := r.currentValue
	ret := r.getCurrentValue()
	if ret != strconv.FormatUint(e, 10) {
		t.Errorf("Got: %v, Expected %v", e, ret)
	}
}
