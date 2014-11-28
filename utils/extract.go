package utils

import (
	"regexp"
	"strconv"
)

// ExtractNum extracts only numeric values out of a string.
func ExtractNum(str string) []string {
	// x|y|z  == match(prefer) x first
	// if we don't delete commas
	// validID := regexp.MustCompile(`(\d{1,3}[,])+\d{1,3}[.]\d{1,}|(\d{1,3}[,])+\d{1,3}|\d{1,}[.]\d{1,}|\d{1,}|[.]\d{1,}`)
	commaID := regexp.MustCompile(`,`)
	str = commaID.ReplaceAllString(str, "")
	validID := regexp.MustCompile(`\d{0,}[.]\d{0,}`)
	numslice := validID.FindAllString(str, -1)
	result := []string{}
	for _, elem := range numslice {
		if elem != "." {
			result = append(result, elem)
		}
	}
	return result
}

// ExtractNumToFloat64 extracts to float64 slice.
func ExtractNumToFloat64(str string) []float64 {
	commaID := regexp.MustCompile(`,`)
	str = commaID.ReplaceAllString(str, "")
	validID := regexp.MustCompile(`\d{0,}[.]\d{0,}`)
	numslice := validID.FindAllString(str, -1)
	// to convert strings into float numbers
	numbers := []float64{}
	for _, elem := range numslice {
		i, err := strconv.ParseFloat(elem, 64)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}
