package bay

import (
	"strings"

	"github.com/gyuho/gobay/slm"
	"github.com/gyuho/goling/st"
)

// GetCdC gets the candidate classs from the training data.
// We assume that the class string is of only one ftw.
func GetCdC(DATA []TD) []int {
	var candidate []int
	for _, elem := range DATA {
		candidate = append(candidate, elem.Class)
	}
	return slm.UniqInts(candidate)
}

// GetCdFt extracts the candidate feature words from the training
// data and feature range data.
// Previous step to mutual information filtering.
// For example, retrieve the useful words: simple, easy, like, hate, etc.
// All raw data are already processed before calling this function.
// This function just extract the raw feature data.
// More informative words will be selected with mutual information.
func GetCdFt(DATA []TD, include string, exclude []string) []string {

	// first candidate
	var rc []string

	for _, each_text := range DATA {
		// for each text, it extracts only words without punctuation
		words := st.GetWords(strings.ToLower(each_text.Text))

		// for each word of current text array elements,
		// test if it is a proper, informative word
		for _, wd := range words {

			// 1st filter: select relatively useful and informative words(signal words)
			// during the 1st filter, the words like "a", "of" will not be caught
			// add more conditions to catch too short words
			// 2nd filter: no exception: exclude useless words for classification
			if strings.Contains(include, strings.ToLower(wd)) && !slm.CheckStr(wd, exclude) {
				rc = append(rc, strings.ToLower(wd))
			}
		}
	}
	return slm.UniqStrsLW(rc)
}
