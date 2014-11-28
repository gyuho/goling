package similar

import "github.com/gyuho/goling/utils"

// Levenshtein is a string metric for measuring the difference between two sequences.
// Levenshtein edit distance between two words is the minimum
// number of single-character edits (insertion, deletion, substitution)
// required to change one word into the other.
// The bigger the return value is, the less similar the two texts are
// because different texts take more edits than similar texts.
// Insertion, deletion, and substitution each costs 1 edit.
func Levenshtein(str1, str2 string) float64 {

	// to clear out the unnecessary characters
	str1 = utils.ReplaceNonAlnum(str1, " ")
	str2 = utils.ReplaceNonAlnum(str2, " ")

	distance := make2Dslice(len(str1)+1, len(str2)+1)

	// initialize the distance array, with position
	for y := 0; y <= len(str1); y++ {
		distance[y][0] = y
	}

	// initialize the distance array, with position
	for x := 0; x <= len(str2); x++ {
		distance[0][x] = x
	}

	for i := 0; i < len(str1); i++ {

		for j := 0; j < len(str2); j++ {

			// edit cost
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

	// equal texts need 0 edit distance
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
