package similar

import "testing"

func TestCosine(t *testing.T) {
	// the bigger the return value is,
	// the more similar the two texts are.

	// very similar texts (0.8616404368553293)
	text1 := "This is a foo bar sentence."
	text2 := "This sentence is similar to a foo bar sentence."
	// very different texts (0.04879500364742665)
	text3 := "AppendFloat appends the string form of the floating-point number f, as generated by FormatFloat, to dst and returns the extended buffer."
	text4 := "There are very simple ways to use matlab in parallel, with GPU, multi-thread support."
	if Cosine(text1, text2) < Cosine(text3, text4) {
		t.Errorf("Error: text1, text2 should be more similar than text3, text4")
	}
	// less different texts (0.31980107453341566)
	text5 := "How is the product team at SEOmoz structured?"
	text6 := "Moz: How trustworthy is SEOMoz after what they did to DOZ?"

	// same texts (1)
	text7 := "I am a gopher."
	text8 := "I am a gopher."
	if Cosine(text5, text6) > Cosine(text7, text8) {
		t.Errorf("Error: text7, text8 should be more simliar than text3, text4")
	}

	// totally different text
	text9 := "I am a gopher."
	text0 := ""
	if Cosine(text9, text0) != 0.0001 {
		t.Errorf("Error: text9, text10 should return 0.0001 with smoothing")
	}
	if Cosine("intention", "execution") != 0.6917144638660746 {
		t.Errorf("Error: intention, execution, Cosine")
	}
	if Cosine("i n t e n t i o n", "e x e c u t i o n") != 0.6917144638660746 {
		t.Errorf("Error: intention, execution, Cosine")
	}
	if Cosine("AGGCTATCACCTGACCTCCAGGCCGATGCC", "TAGCTATCACGACCGCGGTCGATTTGCCCGAC") != 0.9890107762526607 {
		t.Errorf("DNA, Cosine")
	}
	if Cosine("like", "lllikee") != 0.9036961141150639 {
		t.Errorf("like", "lllikee", "Cosine")
	}
	if Cosine("love it", "go to the beach") != 0.0001 {
		t.Errorf("love it", "go to the beach", "Cosine")
	}
	sim00 := Cosine("Spokesman confirms senior government adviser was shot.", "Spokesman said the senior adviser was shot dead.")
	csim00 := 0.5345224838248487
	if sim00 != csim00 {
		t.Errorf("%f != %f (%v)", sim00, csim00, sim00-csim00)
	}
}
