package st

import "testing"

func Test_CountCharacter(t *testing.T) {
	str := "Hello Hi"
	r := CountCharacter(str)

	if r != 8 {
		t.Errorf("CountCharacter(str) should return 8: %#v", r)
	}
}

func Test_CountWord(t *testing.T) {
	str := "Hello World! This is Go."
	r := CountWord(str)

	if r != 5 {
		t.Errorf("CountWord(str) should return 5: %#v", r)
	}
}

func Test_CountSentence(t *testing.T) {
	str := "Hello World! This is Go."
	r := CountSentence(str)

	if r != 2 {
		t.Errorf("CountSentence(str) should return 2: %#v", r)
	}
}
