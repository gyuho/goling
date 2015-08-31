package spellcheck

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
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
	now := time.Now()
	r1 := Suggest("korrecter", fmap)
	fmt.Printf("r1 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r1 != "corrected" {
		t.Errorf("r1 should return 'corrected': %#v", r1)
	}

	now = time.Now()
	r2 := Suggest("hete", fmap)
	fmt.Printf("r2 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r2 != "here" {
		t.Errorf("r2 should return 'here': %#v", r2)
	}

	now = time.Now()
	r3 := Suggest("somthang", fmap)
	fmt.Printf("r3 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r3 != "something" {
		t.Errorf("r3 should return 'something': %#v", r3)
	}

	now = time.Now()
	r4 := Suggest("liike", fmap)
	fmt.Printf("r4 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r4 != "like" {
		t.Errorf("r4 should return 'like': %#v", r4)
	}

	now = time.Now()
	r5 := Suggest("huose", fmap)
	fmt.Printf("r5 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r5 != "house" {
		t.Errorf("r5 should return 'house': %#v", r5)
	}

	now = time.Now()
	r6 := Suggest("tahnks", fmap)
	fmt.Printf("r6 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r6 != "tanks" {
		t.Errorf("r6 should return 'tanks': %#v", r6)
	}

	now = time.Now()
	r7 := Suggest("gool", fmap)
	fmt.Printf("r7 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r7 != "good" {
		t.Errorf("r7 should return 'good': %#v", r7)
	}

	now = time.Now()
	r8 := Suggest("Fransisco", fmap)
	fmt.Printf("r8 takes %v\n", float64(time.Since(now))/float64(1e9))
	if r8 != "Fransisco" {
		t.Errorf("r8 should return 'Fransisco': %#v", r8)
	}
}
