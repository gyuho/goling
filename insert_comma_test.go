package goling

import "testing"

func Test_InsertCommaInt(t *testing.T) {
	var num int64 = 21391239213
	r := InsertCommaInt(num)

	if r != "21,391,239,213" {
		t.Errorf("InsertCommaInt(num) should return \"21,391,239,213\" : %#v", r)
	}
}

func Test_InsertCommaFloat(t *testing.T) {
	var num float64 = 21391239213.5122123
	r := InsertCommaFloat(num)

	if r != "21,391,239,213.51221" {
		t.Errorf("InsertCommaFloat(num) should return \"21,391,239,213.51221\": %#v", r)
	}
}
