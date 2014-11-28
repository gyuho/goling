package utils

import "testing"

func TestExtractNum(t *testing.T) {
	str := "I have $1,500,000. The car is $79,000. But I pay only 20 dollars, and the tax was $10.57. The final price is $80,000.35. BRC Burrito is $0.50 right now. I've got change of $.20."
	r := ExtractNum(str)
	s := []string{"1500000.", "79000.", "10.57", "80000.35", "0.50", ".20"}
	for idx, elem := range s {
		if r[idx] != elem {
			t.Errorf("Not same: \n%#v\n%#v", r, s)
		}
	}
}

func TestExtractNumToFloat64(t *testing.T) {
	str := "I have $1,500,000. The car is $79,000. But I pay only 20 dollars, and the tax was $10.57. The final price is $80,000.35. BRC Burrito is $0.50 right now. I've got change of $.20."
	r := ExtractNumToFloat64(str)
	s := []float64{1500000.000000, 79000.000000, 10.570000, 80000.350000, 0.500000, 0.200000}
	for idx, elem := range s {
		if r[idx] != elem {
			t.Errorf("Not same: \n%#v\n%#v", r, s)
		}
	}
}
