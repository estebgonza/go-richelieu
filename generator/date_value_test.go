package generator

import (
	"reflect"
	"testing"
	"time"
)

func TestDateValueInit1(t *testing.T) {
	var r1 dateValue
	var r2 dateValue
	e := time.Date(
		2009, 11, 17, 20, 34, 58, 000000000, time.UTC)
	initValue := "2009-11-17 20:34:58"
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

func TestDateValueInit2(t *testing.T) {
	var etype time.Time
	var r1 dateValue
	var r2 dateValue
	initValue1 := "2009-11-17 20:34:58"
	initValue2 := "2012-02-01 00:00:00"
	r1.init(initValue1)
	r2.init(initValue2)
	e1 := time.Date(
		2009, 11, 17, 20, 34, 58, 000000000, time.UTC)
	e2 := time.Date(
		2012, 02, 01, 00, 00, 00, 000000000, time.UTC)
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

func TestDateValueGenerate1(t *testing.T) {
	var r dateValue
	initValue := "2009-11-17 20:34:58"
	r.init(initValue)
	e := time.Date(
		2009, 11, 17, 20, 34, 59, 000000000, time.UTC)
	r.generateValue()
	if r.currentValue != e {
		t.Errorf("Expected %v. Got %v", e, r.currentValue)
	}
}

func TestDateValueGetCurrentValue1(t *testing.T) {
	var r dateValue
	initValue := "2009-11-17 20:34:58"
	r.init(initValue)
	var e string = "2009-11-17 20:34:58 +0000 UTC"
	ret := r.getCurrentValue()
	if ret != e {
		t.Errorf("Expected %v. Got %v", e, ret)
	}
}
