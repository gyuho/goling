package spellcheck

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

func TestSuggest(t *testing.T) {
	fmap := Frequency(open("../testdata/sample.txt"))
	r1 := Suggest("korrecter", fmap)
	if r1 != "corrected" {
		t.Errorf("r1 should return 'corrected': %#v", r1)
	}

	r2 := Suggest("hete", fmap)
	if r2 != "here" {
		t.Errorf("r2 should return 'here': %#v", r2)
	}

	r3 := Suggest("somthang", fmap)
	if r3 != "something" {
		t.Errorf("r3 should return 'something': %#v", r3)
	}

	r4 := Suggest("liike", fmap)
	if r4 != "like" {
		t.Errorf("r4 should return 'like': %#v", r4)
	}

	r5 := Suggest("huose", fmap)
	if r5 != "house" {
		t.Errorf("r5 should return 'house': %#v", r5)
	}

	r6 := Suggest("tahnks", fmap)
	if r6 != "tanks" {
		t.Errorf("r6 should return 'tanks': %#v", r6)
	}

	r7 := Suggest("gool", fmap)
	if r7 != "good" {
		t.Errorf("r7 should return 'good': %#v", r7)
	}

	r8 := Suggest("Fransisco", fmap)
	if r8 != "Fransisco" {
		t.Errorf("r8 should return 'Fransisco': %#v", r8)
	}
}
