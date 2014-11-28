package slm

import (
	"reflect"
	"testing"
)

func Test_MaxIntFK(test *testing.T) {
	m := map[float64]int{
		0.9:  9,
		0.1:  1,
		3.14: 314,
		5.1:  51,
		5.0:  51,
		1.1:  51,
	}
	if MaxIntFK(m) != 51 {
		test.Errorf("It should return 51 but %v", MaxIntFK(m))
	}
}

func Test_Top5SFK(test *testing.T) {
	m := map[float64]string{
		0.9:  "9",
		0.1:  "1",
		3.14: "314",
		5.1:  "51",
		5.0:  "51",
		1.1:  "51",
	}
	if CheckStr("1", Top5SFK(m)) {
		test.Errorf("1 shouldn't be in the result: %v", Top5SFK(m))
	}
}

func Test_UniqInts(test *testing.T) {
	s := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	if len(UniqInts(s)) != 5 {
		test.Errorf("Should return 5 but %v", UniqInts(s))
	}
}

func Test_StrToInt(test *testing.T) {
	if reflect.TypeOf(StrToInt("10")) != reflect.TypeOf(10) {
		test.Errorf("Should return true but %v", reflect.TypeOf(StrToInt("10")))
	}
}

func Test_IsecStrsLC(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	b := []string{"C", "D", "E", "F"}
	if len(IsecStrsLC(a, b)) != 2 {
		test.Errorf("Should return 2 but %v", IsecStrsLC(a, b))
	}
}

func Test_CheckStr(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	if CheckStr("A", a) != true {
		test.Errorf("Should return true but %v", CheckStr("A", a))
	}
}

func Test_SubStrsLC(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	b := []string{"C", "D", "E", "F"}
	if len(SubStrsLC(a, b)) != 2 {
		test.Errorf("Should return 2 but %v", SubStrsLC(a, b))
	}
}

func Test_UnionStrsLC(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	b := []string{"C", "D", "E", "F"}
	if len(UnionStrsLC(a, b)) != 8 {
		test.Errorf("Should return 8 but %v", UnionStrsLC(a, b))
	}
}

func Test_DupStrsLW(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	b := []string{"C", "D", "E", "F"}
	if len(DupStrsLW(UnionStrsLC(a, b))) != 2 {
		test.Errorf("Should return 8 but %v", DupStrsLW(UnionStrsLC(a, b)))
	}
}

func Test_UniqStrsLW(test *testing.T) {
	a := []string{"A", "B", "C", "D"}
	b := []string{"C", "D", "E", "F"}
	if len(UniqStrsLW(UnionStrsLC(a, b))) != 6 {
		test.Errorf("Should return 8 but %v", UniqStrsLW(UnionStrsLC(a, b)))
	}
}
