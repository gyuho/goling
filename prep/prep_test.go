package prep

import (
	"strings"
	"testing"
)

func TestCleanUp(t *testing.T) {
	str1 := "   Hello,    World! 124  2 This 23is Go project,		It is amazing  . I am excited!    Are you too?    "
	r1 := CleanUp(str1)
	if r1 != "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?" {
		t.Errorf("Expected \"Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?\": %s", r1)
	}
	str2 := "		Hello World!	This is Go 		"
	r2 := CleanUp(str2)
	if r2 != "Hello World! This is Go" {
		t.Errorf("Expected \"Hello World! This is Go\": %s", r2)
	}
	str3 := "	\n\n\n	Hello World! \n\n	This is Go 		"
	r3 := CleanUp(str3)
	if r3 != "Hello World! This is Go" {
		t.Errorf("Expected \"Hello World! This is Go\": %s", r3)
	}
}

func TestExpand(t *testing.T) {
	str := "was't weren't aren't isn't I'd you'd I'll he doesn't Don't I won't I'm You're i've what's up it's I didn't"
	r := Expand(strings.ToLower(str))
	rc := "was not were not are not is not i would you would i will he does not do not i will not i am you are i have what is up it is i did not"
	if r != rc {
		t.Errorf("Expected\n%v\nbut\n%v", rc, r)
	}
}

func TestReplaceNonAlpha(t *testing.T) {
	str1 := "Hello 1231thi21s123 is 12G33o25!!!!"
	r1 := ReplaceNonAlpha(strings.ToLower(str1), " ")
	if r1 != "hello     thi  s    is   g  o      " {
		t.Errorf("Expected \"hello     thi  s    is   g  o      \": %s", r1)
	}
	str2 := "Hello 1231thi21s123 is 12G33o25!!!!"
	r2 := ReplaceNonAlpha(strings.ToLower(str2), "")
	if r2 != "hello this is go" {
		t.Errorf("Expected \"hello this is go\": %s", r2)
	}
}

func TestReplaceNonAlnum(t *testing.T) {
	str1 := "It's me! Hello !@##%W@!#orl!@#!@#!@#d!"
	r1 := ReplaceNonAlnum(Expand(strings.ToLower(str1)), " ")
	if r1 != "it is me  hello      w   orl         d " {
		t.Errorf("Expected \"it is me  hello      w   orl         d \": %s", r1)
	}
	str2 := "Hello !@##%W@!#orl!@#!@#!@#d!"
	r2 := ReplaceNonAlnum(Expand(strings.ToLower(str2)), "")
	if r2 != "hello world" {
		t.Errorf("Expected \"hello world\": %s", r2)
	}
}

func TestSplitToWords(t *testing.T) {
	slice1 := SplitToWords("sadf \nsdafsdf a  s  \nsadfsdf")
	if len(slice1) != 5 {
		t.Errorf("Expected 5 but %#v", slice1)
	}
	str2 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	slice2 := SplitToWords(strings.ToLower(str2))
	cslice2 := []string{"hello", "world!", "it", "is", "good", "to", "see", "you,", "them,", "10.", "this", "is", "go.", "how", "are", "you?", "he", "said", "i", "am", "x1000", "good."}
	if len(slice2) != len(cslice2) {
		t.Errorf("%#v\n%#v", slice2, cslice2)
	}
	for k, v := range slice2 {
		if cslice2[k] != v {
			t.Errorf("%#v != %#v", cslice2[k], v)
		}
	}
}

func TestSplitToSentences(t *testing.T) {
	str1 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	slice1 := SplitToSentences(str1)
	cslice1 := []string{"Hello World!", "It is Good to See You, Them, 10.", "This is Go.", "How are you?", "He said I am x1000 good."}
	if len(slice1) != len(cslice1) {
		t.Errorf("%#v\n%%#v", slice1, cslice1)
	}
	for k, v := range slice1 {
		if cslice1[k] != v {
			t.Errorf("%#v != %#v", cslice1[k], v)
		}
	}
	str2 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
	slice2 := SplitToSentences(str2)
	cslice2 := []string{"Hello World!", "It is Good to See You, Them, 10.", "This is Go.", "How are you?", "He said I am x1000 good."}
	if len(slice2) != len(cslice2) {
		t.Errorf("%#v\n%%#v", slice2, cslice2)
	}
	for k, v := range slice2 {
		if cslice2[k] != v {
			t.Errorf("%#v != %#v", cslice2[k], v)
		}
	}
}

func TestCountCharacter(t *testing.T) {
	str := "Hello Hi"
	r := CountCharacter(str)
	if r != 8 {
		t.Errorf("CountCharacter(str) should return 8: %#v", r)
	}
}

func TestCountWord(t *testing.T) {
	str := "Hello World! This is Go."
	r := CountWord(str)
	if r != 5 {
		t.Errorf("CountWord(str) should return 5: %#v", r)
	}
}

func TestCountSentence(t *testing.T) {
	str := "Hello World! This is Go."
	r := CountSentence(str)
	if r != 2 {
		t.Errorf("CountSentence(str) should return 2: %#v", r)
	}
}
