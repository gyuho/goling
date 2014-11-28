package utils

import "testing"

func TestInsertCommaInt64(t *testing.T) {
	var num int64 = 21391239213
	r := InsertCommaInt64(num)

	if r != "21,391,239,213" {
		t.Errorf("InsertCommaInt(num) should return \"21,391,239,213\" : %#v", r)
	}
}

func TestInsertCommaFloat64(t *testing.T) {
	// var num float64 = 21391239213.5122123
	num := 21391239213.5122123
	r := InsertCommaFloat64(num)

	if r != "21,391,239,213.51221" {
		t.Errorf("InsertCommaFloat(num) should return \"21,391,239,213.51221\": %#v", r)
	}
}
