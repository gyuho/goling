package st

import (
	"regexp"
	"strconv"
)

// ExtractNum extracts only numeric values out of a string.
func ExtractNum(str string) []float64 {

	// x|y|z  == match(prefer) x first
	// if we don't delete commas
	// validID := regexp.MustCompile(`(\d{1,3}[,])+\d{1,3}[.]\d{1,}|(\d{1,3}[,])+\d{1,3}|\d{1,}[.]\d{1,}|\d{1,}|[.]\d{1,}`)

	commaID := regexp.MustCompile(`,`)
	str = commaID.ReplaceAllString(str, "")

	validID := regexp.MustCompile(`\d{0,}[.]\d{0,}`)
	numslice := validID.FindAllString(str, -1)

	// to convert strings into float numbers
	var numbers []float64
	for _, elem := range numslice {
		i, err := strconv.ParseFloat(elem, 64)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

// ExtractNumAndSum extracts the numeric values and add up the numbers.
func ExtractNumAndSum(str string) float64 {
	numbers := ExtractNum(str)
	total := 0.0
	for _, elem := range numbers {
		total += elem
	}
	return total
}
