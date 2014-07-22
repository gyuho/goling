package simi

import (
	"math"
	"strings"

	"github.com/gyuho/goling/st"
)

// CosineSimilarity is word-wise.
// LevenshteinDistance is character-wise.
// Both are useful to calculate the similarities of texts.
// But note that this does not consider any meanings in context.
// They literally calculate the similarities in physical letters.
// CosineSimilarity is a measure of similarity between two vectors.
// Return the cosine similarity of two texts.
// similarity = cos(θ) = A·B/|A|*|B|
// Used for analyzing the string similarity.
// The similarity scales between 0 and 1(maximum).
// The bigger the return value is, the more similar the two texts are.
// cos 0° = 1, cos 90° = 0
// 0° means that the two texts are equal, since two sequences point to the same point.
// 90° means that the two texts are totally different.
func CosineSimilarity(str1, str2 string) float64 {
	// to clear out the unnecessary characters
	// And consider the case that the string has only one word.
	temp1 := st.GetWords(str1)
	sli1 := []string{}
	if len(temp1) == 1 {
		// func Split(s, sep string) []string
		sli1 = strings.Split(temp1[0], "")
	} else {
		sli1 = temp1
	}

	temp2 := st.GetWords(str2)
	sli2 := []string{}
	if len(temp2) == 1 {
		// func Split(s, sep string) []string
		sli2 = strings.Split(temp2[0], "")
	} else {
		sli2 = temp2
	}

	// to convert string to vector
	// to associate each character with its frequency values in string

	// vect1 contains the frequencies of each character
	vect1 := make(map[string]int)
	for _, elem := range sli1 {
		vect1[elem] += 1
	}

	// vect2 contains the frequencies of each character
	vect2 := make(map[string]int)
	for _, elem := range sli2 {
		vect2[elem] += 1
	}

	// intersection contains common characters
	intersection := []string{}

	// traverse keys and add what is common to both.
	// (keys, in Hash/Map, are like indices in array)
	for key := range vect1 {
		if _, exist := vect2[key]; exist {
			intersection = append(intersection, key)
		}
	}

	// If all the vector elements are equal, cos will be 1.
	// Equal texts return the value 1.
	// We need to traverse the intersection(common) characters of texts.
	// In doing so, we can expect two same texts to return 1, cos 0°

	// to calculate A·B
	sum := 0.0
	for _, elem := range intersection {
		sum += float64(vect1[elem]) * float64(vect2[elem])
	}
	numerator := sum

	// to calculate |A|*|B|
	sum1 := 0.0
	for key := range vect1 {
		sum1 += math.Pow(float64(vect1[key]), 2)
	}

	sum2 := 0.0
	for key := range vect2 {
		sum2 += math.Pow(float64(vect2[key]), 2)
	}
	denominator := math.Sqrt(sum1) * math.Sqrt(sum2)

	// because we can't divide by 0
	if numerator == 0.0 || denominator == 0.0 {
		return 0.0001
	} else {
		return float64(numerator) / denominator
	}
}
