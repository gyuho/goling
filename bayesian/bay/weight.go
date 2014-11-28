package bay

import "strings"

// TotalWt returns the total weight value.
func TotalWt(DATA []TD) int {
	total := 0
	for _, elem := range DATA {
		total += elem.Weight
	}
	return total
}

// WtByF returns the total weight value of certain feature.
func WtByF(DATA []TD, ftw string) int {
	total := 0
	for _, elem := range DATA {
		if strings.Contains(strings.ToLower(elem.Text), strings.ToLower(ftw)) {
			total += elem.Weight
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// WtByNF returns the total weight value when the input feature does not occur.
// W(~"like")
func WtByNF(DATA []TD, ftw string) int {
	total := 0
	for _, elem := range DATA {
		if !strings.Contains(strings.ToLower(elem.Text), strings.ToLower(ftw)) {
			total += elem.Weight
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// WtByC returns the total weight value of certain class.
func WtByC(DATA []TD, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.Class == klass {
			total += elem.Weight
		}
	}
	return total
}

// WtByFC returns the total weight value by both class and feature words.
// For example, use this to get "like" in "positive" class.
func WtByFC(DATA []TD, ftw string, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.Class == klass {
			if strings.Contains(strings.ToLower(elem.Text), strings.ToLower(ftw)) {
				total += elem.Weight
			}
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// WtByNFC returns the total weight value by class and with the Non-feature word.
// For example, use this to get "like" in "positive" class.
func WtByNFC(DATA []TD, ftw string, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.Class == klass {
			if !strings.Contains(strings.ToLower(elem.Text), strings.ToLower(ftw)) {
				total += elem.Weight
			}
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}
