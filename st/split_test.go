package st

import "testing"

func Test_SplitWord(t *testing.T) {
	str := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	r := SplitWord(str)
	a := []string{"hello", "world!", "it", "is", "good", "to", "see", "you,", "them,", "10.", "this", "is", "go.", "how", "are", "you?", "he", "said", "i", "am", "x1000", "good."}
	if len(r) != len(a) {
		t.Errorf("SplitWord(str) should return the same length: %#v", r)
	}
	for k, v := range r {
		if a[k] != v {
			t.Errorf("a[k] != v: %#v", a[k])
		}
	}
}

func Test_GetWords(t *testing.T) {
	str := "Hello Hi"
	r := GetWords(str)
	a := []string{"hello", "hi"}
	if len(r) != len(a) {
		t.Errorf("GetWords(str) should return the same length: %#v", r)
	}
	for k, v := range r {
		if a[k] != v {
			t.Errorf("a[k] != v: %#v", a[k])
		}
	}
}

func Test_SplitSentRaw(t *testing.T) {
	str := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	r := SplitSentRaw(str)
	a := []string{"Hello World!", "It is Good to See You, Them, 10.", "This is Go.", "How are you?", "He said I am x1000 good."}
	if len(r) != len(a) {
		t.Errorf("SplitSentRaw(str) should return the same length: %#v", r)
	}
	for k, v := range r {
		if a[k] != v {
			t.Errorf("a[k] != v: %#v", a[k])
		}
	}
}

func Test_SplitSent(t *testing.T) {
	str := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	r := SplitSent(str)
	a := []string{"Hello World", "It is Good to See You, Them, 10", "This is Go", "How are you", "He said I am x1000 good"}
	if len(r) != len(a) {
		t.Errorf("SplitSent(str) should return the same length: %#v", r)
	}
	for k, v := range r {
		if a[k] != v {
			t.Errorf("a[k] != v: %#v", a[k])
		}
	}
}

func Test_SplitParag(t *testing.T) {
	str := "Hello World! \n This is Go!"
	r := SplitParag(str)
	a := []string{"Hello World!", "This is Go!"}
	if len(r) != len(a) {
		t.Errorf("SplitParag(str) should return the same length: %#v", r)
	}
	for k, v := range r {
		if a[k] != v {
			t.Errorf("a[k] != v: %#v", a[k])
		}
	}
}
