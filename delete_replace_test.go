package goling

import "testing"

func Test_ReplaceNonAlphaWithSpace(t *testing.T) {
	str := "Hello 1231thi21s123 is 12G33o25!!!!"
	r := ReplaceNonAlphaWithSpace(str)

	if r != "hello     thi  s    is   g  o      " {
		t.Errorf("ReplaceNonAlphaWithSpace(str) should return \"hello     thi  s    is   g  o      \": %#v", r)
	}
}

func Test_DeleteNonAlpha(t *testing.T) {
	str := "Hello 1231thi21s123 is 12G33o25!!!!"
	r := DeleteNonAlpha(str)

	if r != "hello this is go" {
		t.Errorf("DeleteNonAlpha(str) should return \"hello this is go\": %#v", r)
	}
}

func Test_ReplaceNonAlnumWithSpace(t *testing.T) {
	str := "It's me! Hello !@##%W@!#orl!@#!@#!@#d!"
	r := ReplaceNonAlnumWithSpace(str)

	if r != "it is me  hello      w   orl         d " {
		t.Errorf("ReplaceNonAlnumWithSpace(str) should return \"it is me  hello      w   orl         d \": %#v", r)
	}
}

func Test_DeleteNonAlnum(t *testing.T) {
	str := "Hello !@##%W@!#orl!@#!@#!@#d!"
	r := DeleteNonAlnum(str)

	if r != "hello world" {
		t.Errorf("DeleteNonAlnum(str) should return \"hello world\": %#v", r)
	}
}
