package segment

import (
	"log"
	"os"
	"testing"
)

func open(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %+v", err)
	}
	return f
}

func TestGet(t *testing.T) {
	// It really depends on the training data.
	// The performance can significantly improve with more training data.

	pFunc1 := Frequency(open("../testdata/sample.txt"))
	r1 := Get("IamCodingwithGO.Ireallyenjoythislanguage.", pFunc1)
	// [IamCodingwithGO.I really enjoy this language .]
	if r1[1] != "really" {
		t.Errorf("r1[1] should return 'really': %#v ", r1[1])
	}

	pFunc2 := Frequency(open("../testdata/sample.txt"))
	r2 := Get("DoYouknowhereiamcoding?Iamcodinginlibrary.", pFunc2)
	// [DoYouknowhereiamcoding?Iamcodinginlibrary.]
	if len(r2) != 1 {
		t.Errorf("len(r2) should return 1: %#v ", r2[1])
	}

	pFunc3 := Frequency(open("../testdata/sample.txt"))
	r3 := Get("TheyaregreatpeopleHeisagreatperson.", pFunc3)
	// [They are great people He is a great person .]
	if r3[1] != "are" {
		t.Errorf("r3[1] should return 'are': %#v ", r3[1])
	}
}
