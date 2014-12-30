package similar // import "github.com/gyuho/goling/similar"

import (
	"math"
	"strings"

	"github.com/gyuho/stringx"
)

// Cosine is word-wise.
// LevenshteinDistance is character-wise.
// Both are useful to calculate the similarities of texts.
// But note that this does not consider any meanings in context.
// They literally calculate the similarities in physical letters.
// Cosine Similarity is a measure of similarity between two vectors.
// Return the cosine similarity of two texts.
// similarity = cos(θ) = A·B/|A|*|B|
// Used for analyzing the string similarity.
// The similarity scales between 0 and 1(maximum).
// The bigger the return value is, the more similar the two texts are.
// cos 0° = 1, cos 90° = 0
// 0° means that the two texts are equal, since two sequences point to the same point.
// 90° means that the two texts are totally different.
func Cosine(str1, str2 string) float64 {
	// to clear out the unnecessary characters
	// And consider the case that the string has only one word.
	temp1 := stringx.SplitToWords(strings.ToLower(str1))
	sli1 := []string{}
	if len(temp1) == 1 {
		sli1 = strings.Split(temp1[0], "")
	} else {
		sli1 = temp1
	}
	temp2 := stringx.SplitToWords(strings.ToLower(str2))
	sli2 := []string{}
	if len(temp2) == 1 {
		sli2 = strings.Split(temp2[0], "")
	} else {
		sli2 = temp2
	}

	// to convert string to vector
	// to associate each character with its frequency values in string
	// vect1 contains the frequencies of each character
	vect1 := make(map[string]int)
	for _, elem := range sli1 {
		vect1[elem]++
	}
	// vect2 contains the frequencies of each character
	vect2 := make(map[string]int)
	for _, elem := range sli2 {
		vect2[elem]++
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

	// smoothing because we can't divide by 0
	if numerator == 0.0 || denominator == 0.0 {
		return 0.0001
	}
	return float64(numerator) / denominator
}

// Levenshtein is a string metric for measuring the difference between two sequences.
// Levenshtein edit distance between two words is the minimum
// number of single-character edits (insertion, deletion, substitution)
// required to change one word into the other.
// The bigger the return value is, the less similar the two texts are
// because different texts take more edits than similar texts.
// Insertion, deletion, and substitution each costs 1 edit.
func Levenshtein(str1, str2 string) float64 {
	// to clear out the unnecessary characters
	str1 = stringx.ReplaceNonAlnum(str1, " ")
	str2 = stringx.ReplaceNonAlnum(str2, " ")

	// initialize the distance array, with position
	distance := make2Dslice(len(str1)+1, len(str2)+1)
	for y := 0; y <= len(str1); y++ {
		distance[y][0] = y
	}
	for x := 0; x <= len(str2); x++ {
		distance[0][x] = x
	}

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			edit := 1
			if str1[i] == str2[j] {
				edit = 0
			}
			// distance[i][j+1] + 1 : delete/insert from str1
			// distance[i+1][j] + 1 : delete/insert from str2
			// distance[i][j] + edit : delete/insert from both
			distance[i+1][j+1] = getMin(
				distance[i][j+1]+1,
				distance[i+1][j]+1,
				distance[i][j]+edit,
			)
		}
	}
	if distance[len(str1)][len(str2)] == 0 {
		// smoothing
		return 0.0001
	}

	return float64(distance[len(str1)][len(str2)])
}

func getMin(more ...int) int {
	min := more[0]
	for _, elem := range more {
		if min > elem {
			min = elem
		}
	}
	return min
}

func make2Dslice(row, column int) [][]int {
	mat := make([][]int, row)
	for i := range mat {
		mat[i] = make([]int, column)
	}
	return mat
}

// Get returns the similarity of two string using Cosing similarity and Levenshtein distance.
// Higher number means higher similarity. It is up to users to preprossess input texts.
func Get(str1, str2 string) float64 {
	// Higher means higher similarity
	sim1 := Cosine(str1, str2)

	// Lower means higher similarity
	siml2 := Levenshtein(str1, str2)
	sim2 := (1.0 / float64(siml2)) * 1000

	siml3 := Levenshtein(stringx.DeleteSpace(str1), stringx.DeleteSpace(str2))
	sim3 := (1.0 / float64(siml3)) * 1000

	return math.Sqrt(sim1*sim2*sim3) * 100
}
