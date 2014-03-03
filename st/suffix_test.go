package st

import "testing"

func Test_DeleteSuffix(t *testing.T) {
	str1 := "Generalization"
	if "general" != DeleteSuffix(str1) {
		t.Errorf("Should return 'general' but %v", DeleteSuffix(str1))
	}
}
