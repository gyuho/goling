package bay

import (
	"math"
	"sort"

	"github.com/gyuho/gobay/slm"
)

// MutInfByFt calculates the mutual information probability
// to detect mutually informative features.
// For example, it returns higher probability for "like"
// rather than "the."
func MutInfByFt(DATA []TD, CLASSES []int, ftw string) float64 {
	result := 0.0
	for _, elem := range CLASSES {
		// P(“hate” ∩ ✙)·log[ P(“hate” ∩ ✙)/{P(“hate”)·P(✙)} ]
		// P(“hate” ∩ -)·log[ P(“hate” ∩ -)/{P(“hate”)·P(-)} ]
		result += JtProbFC(DATA, ftw, elem) * math.Log10(JtProbFC(DATA, ftw, elem)/(ProbByF(DATA, ftw)*ProbByC(DATA, elem)))
		// + P(~“hate” ∩ ✙) · log [  P(~“hate” ∩ ✙) / {P(~“hate”)·P(✙)} ]
		// + P(~“hate” ∩ -) · log [  P(~“hate” ∩ -) / {P(~“hate”)·P(-)} ]
		result += JtProbNFC(DATA, ftw, elem) * math.Log10(JtProbNFC(DATA, ftw, elem)/(ProbByNF(DATA, ftw)*ProbByC(DATA, elem)))
	}
	return result
}

// InfTop5Ft, from mutual information, extracts the most informative features.
func InfTop5Ft(DATA []TD, CLASSES []int, fts []string) []string {
	mtim := make(map[float64]string)

	for _, elem := range fts {
		mtim[MutInfByFt(DATA, CLASSES, elem)] = elem
	}

	return slm.Top5SFK(mtim)
}

// Return the most informative n words.
func InfFtWd(DATA []TD, CLASSES []int, fts []string, howmany int) []string {
	mtim := make(map[float64]string)

	for _, elem := range fts {
		mtim[MutInfByFt(DATA, CLASSES, elem)] = elem
	}

	var keys []float64
	for k := range mtim {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the last index has the biggest float key

	var result []string

	if len(keys) < howmany {
		howmany = len(keys)
	}
	for i := 1; i < howmany; i++ {
		result = append(result, mtim[keys[len(keys)-i]])
	}

	return result
}
