package goling

import "testing"

func Test_ExtractNum(t *testing.T) {
	str := "I have $1,500,000. The car is $79,000. But I pay only 20 dollars, and the tax was $10.57. The final price is $80,000.35. BRC Burrito is $0.50 right now. I've got change of $.20."
	r := ExtractNum(str)
	s := []float64{1500000.000000, 79000.000000, 10.570000, 80000.350000, 0.500000, 0.200000}
	if r[0] != s[0] {
		t.Errorf("r[0] is not same as s[0]: %#v", r)
	}

	if r[1] != s[1] {
		t.Errorf("r[1] is not same as s[1]: %#v", r)
	}

	if r[2] != s[2] {
		t.Errorf("r[2] is not same as s[2]: %#v", r)
	}

	if r[3] != s[3] {
		t.Errorf("r[3] is not same as s[3]: %#v", r)
	}

	if r[4] != s[4] {
		t.Errorf("r[4] is not same as s[4]: %#v", r)
	}

	if r[5] != s[5] {
		t.Errorf("r[5] is not same as s[5]: %#v", r)
	}
}

func Test_ExtractNumAndSum(t *testing.T) {
	str := "I have $1,500,000. The car is $79,000. But I pay only 20 dollars, and the tax was $10.57. The final price is $80,000.35. BRC Burrito is $0.50 right now. I've got change of $.20."
	r := ExtractNumAndSum(str)

	if r != 1659011.620000 {
		t.Errorf("ExtractNumAndSum(str) should return 1659011.620000: %#v", r)
	}
}
