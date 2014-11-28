// Package slm performs slice and map operations for gobay.
package slm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// MaxIntFK returns the value with the maximum float key.
func MaxIntFK(m map[float64]int) int {
	var keys []float64
	// traverse map only with keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the input map m is sorted by the keys

	return m[keys[len(keys)-1]]
}

// Top5SFK returns top five strings with maximum float64 keys.
func Top5SFK(m map[float64]string) []string {
	var keys []float64
	// traverse map only with keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the input map m is sorted by the keys

	// optional statement
	// fmt.Println("The value with the biggest key is"
	// , m[keys[len(keys)-1]], ", with its key", keys[len(keys)-1])

	var result []string
	result = append(result, m[keys[len(keys)-1]])
	result = append(result, m[keys[len(keys)-2]])
	result = append(result, m[keys[len(keys)-3]])
	result = append(result, m[keys[len(keys)-4]])
	result = append(result, m[keys[len(keys)-5]])

	return result
}

// UniqInts deletes the duplicate elements
// and returns the new slice of unique elements.
func UniqInts(slice []int) []int {
	model := make(map[int]bool)
	result := []int{}

	// traverse the input array and map each to boolean value
	for _, v := range slice {
		if _, checked := model[v]; !checked {
			result = append(result, v)
			model[v] = true
		}
	}
	return result
}

// StrToInt converts string to integer.
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		panic("Fail")
	}

	return i
}

// IsecStrsLC returns the intersection of input slices.
func IsecStrsLC(sl1 []string, sl2 []string) []string {
	return DupStrsLW(UnionStrsLC(UniqStrsLW(sl1), UniqStrsLW(sl2)))
}

// CheckStr checks in lowercse, if the string
// exists in the slice.
func CheckStr(str string, slice []string) bool {
	for _, v := range slice {
		if strings.ToLower(str) == strings.ToLower(v) {
			return true
		}
	}
	return false
}

// SubStrsLC subtracts B from A: A - B
func SubStrsLC(a, b []string) []string {
	result := []string{}
	intersect := IsecStrsLC(a, b)

	for _, v := range a {
		if !CheckStr(v, intersect) {
			result = append(result, v)
		}
	}
	return result
}

// UnionStrsLC combines string slices; duplication is allowed.
// It convert to lower case.
func UnionStrsLC(sl1 []string, sl2 []string) []string {
	// convert all elements to lower case
	ns1 := []string{}
	for _, v := range sl1 {
		ns1 = append(ns1, strings.ToLower(v))
	}
	ns2 := []string{}
	for _, v := range sl2 {
		ns2 = append(ns2, strings.ToLower(v))
	}

	var total []string
	total = append(total, ns1...)
	total = append(total, ns2...)
	// (X) total = append(total, sl1..., sl2...)
	return total
}

// DupStrsLW returns the duplicate elements in slice.
// It convert to lower case.
func DupStrsLW(slice []string) []string {
	// convert all elements to lower case
	ns := []string{}
	for _, v := range slice {
		ns = append(ns, strings.ToLower(v))
	}

	freq := make(map[string]int)
	result := []string{}

	for _, v := range ns {
		freq[v] += 1
		if freq[v] >= 2 {
			result = append(result, v)
		}
	}
	return result
}

// UniqStrsLW deletes the duplicate elements
// and returns the new string slice of unique elements.
// It convert to lower case.
func UniqStrsLW(slice []string) []string {
	// convert all elements to lower case
	ns := []string{}
	for _, v := range slice {
		ns = append(ns, strings.ToLower(v))
	}

	// var model map[string]bool
	// model := map[string]bool{}
	model := make(map[string]bool)

	// var result []string
	// result := make([]string, 5)
	result := []string{}

	// traverse the input array and map each to boolean value
	for _, v := range ns {
		if _, checked := model[v]; !checked {
			result = append(result, v)
			model[v] = true
		}
	}
	return result
}
