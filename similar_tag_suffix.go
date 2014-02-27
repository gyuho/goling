package goling

import (
	"sort"
	"strings"
)

// Return the similarity between two tag arrays.
// The higher, the more similar.
func TagSimilarity(slice1, slice2 []string) float64 {
	sort.Sort(sort.StringSlice(slice1))
	sort.Sort(sort.StringSlice(slice2))

	// fmt.Println(slice1)
	// fmt.Println(slice2)

	nstr1 := strings.Join(slice1, " ")
	nstr1 = ExpandApostrophe(nstr1)
	nstr1 = ReplaceNonAlphaWithSpace(nstr1)

	nstr2 := strings.Join(slice2, " ")
	nstr2 = ExpandApostrophe(nstr2)
	nstr2 = ReplaceNonAlphaWithSpace(nstr2)

	return StringSimilarity(nstr1, nstr2)
}
