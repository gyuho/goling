package seg

import "testing"

func Test_GetSegment(t *testing.T) {
	// It really depends on the training data.
	// The performance can significantly improve with more training data.

	wordp1 := ConstructWordProb(OpenFile("../data/traindata_Dixon.txt"))
	r1 := GetSegment("IamCodingwithGO.Ireallyenjoythislanguage.", wordp1)
	// [IamCodingwithGO.I really enjoy this language .]
	if r1[1] != "really" {
		t.Errorf("r1[1] should return 'really': %#v ", r1[1])
	}

	wordp2 := ConstructWordProb(OpenFile("../data/traindata_Dixon.txt"))
	r2 := GetSegment("DoYouknowhereiamcoding?Iamcodinginlibrary.", wordp2)
	// [DoYouknowhereiamcoding?Iamcodinginlibrary.]
	if len(r2) != 1 {
		t.Errorf("len(r2) should return 1: %#v ", r2[1])
	}

	wordp3 := ConstructWordProb(OpenFile("../data/traindata_Dixon.txt"))
	r3 := GetSegment("TheyaregreatpeopleHeisagreatperson.", wordp3)
	// [They are great people He is a great person .]
	if r3[1] != "are" {
		t.Errorf("r3[1] should return 'are': %#v ", r3[1])
	}
}
