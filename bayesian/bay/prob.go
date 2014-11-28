package bay

// ProbByF returns the probability of class in total cases.
// P(Feature)
func ProbByF(DATA []TD, ftw string) float64 {
	return float64(WtByF(DATA, ftw)) / float64(TotalWt(DATA))
}

// ProbByNF returns the probability of feature NOT occurring.
// P(~Feature)
func ProbByNF(DATA []TD, ftw string) float64 {
	return float64(WtByNF(DATA, ftw)) / float64(TotalWt(DATA))
}

// ProbByC returns the probability of class in total cases.
// P(Class)
func ProbByC(DATA []TD, klass int) float64 {
	return float64(WtByC(DATA, klass)) / float64(TotalWt(DATA))
}

// ProbByFC returns the conditional probaility between feature and class.
// P(Feature | Class)
// For example, use this to get P("like"|+)
func ProbByFC(DATA []TD, ftw string, klass int) float64 {
	return float64(WtByFC(DATA, ftw, klass)) / float64(WtByC(DATA, klass))
}

// JtProbFC returns the joint probability of feature and class.
// P(Feature ∩ Class)
func JtProbFC(DATA []TD, ftw string, klass int) float64 {
	return float64(WtByFC(DATA, ftw, klass)) / float64(TotalWt(DATA))
}

// JtProbNFC returns the joint probability of Non-feature and class.
// P(Feature ∩ Class)
func JtProbNFC(DATA []TD, ftw string, klass int) float64 {
	return float64(WtByNFC(DATA, ftw, klass)) / float64(TotalWt(DATA))
}
