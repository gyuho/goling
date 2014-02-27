package goling

// LevenshteinDistance, edit distance, is a string metric
// for measuring the difference between two sequences.
// Levenshtein distance between two words is the minimum
// number of single-character edits (insertion, deletion, substitution)
// required to change one word into the other.
// The bigger the return value is,
// the less similar the two texts are
// because different texts take more edits than similar texts.
// Insertion, deletion, and substitution each costs 1 edit.
func LevenshteinDistance(str1, str2 string) float64 {

	// to clear out the unnecessary characters
	nstr1 := ReplaceNonAlnumWithSpace(str1)
	nstr2 := ReplaceNonAlnumWithSpace(str2)

	distance := Make2DSlice(len(nstr1)+1, len(nstr2)+1)

	// initialize the distance array, with position
	for y := 0; y <= len(nstr1); y++ {
		distance[y][0] = y
	}

	// initialize the distance array, with position
	for x := 0; x <= len(nstr2); x++ {
		distance[0][x] = x
	}

	for i := 0; i < len(nstr1); i++ {

		for j := 0; j < len(nstr2); j++ {

			// edit cost
			edit := 1
			if nstr1[i] == nstr2[j] {
				edit = 0
			}

			// distance[i][j+1] + 1 : delete/insert from nstr1
			// distance[i+1][j] + 1 : delete/insert from nstr2
			// distance[i][j] + edit : delete/insert from both
			distance[i+1][j+1] = GetMinIntMore(distance[i][j+1]+1,
				distance[i+1][j]+1,
				distance[i][j]+edit)
		}
	}

	// equal texts need 0 edit distance
	if distance[len(nstr1)][len(nstr2)] == 0 {
		// smoothing
		return 0.0001
	} else {
		return float64(distance[len(nstr1)][len(nstr2)])
	}
}

// GetMinIntMore returns the minimum integer.
func GetMinIntMore(more ...int) int {
	min := more[0]
	for _, elem := range more {
		if min > elem {
			min = elem
		}
	}
	return min
}

func Make2DSlice(row, column int) [][]int {
	mat := make([][]int, row)
	// for i := 0; i < row; i++ {
	for i := range mat {
		mat[i] = make([]int, column)
	}
	return mat
}
