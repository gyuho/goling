package st

import "testing"

func Test_Stem(t *testing.T) {
	str1 := "Generalization"
	if "general" != Stem(str1) {
		t.Errorf("Should return 'general' but %v", Stem(str1))
	}
}

func Test_StemText(t *testing.T) {
	str1 := "He is doing Generalization"
	if "he is do general" != StemText(str1) {
		t.Errorf("Should return 'he is do general' but %v", StemText(str1))
	}
}
