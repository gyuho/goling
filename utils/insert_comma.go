package utils

import (
	"strconv"
	"strings"
)

// InsertCommaInt64 inserts comma in every three digit.
// It returns the new version of input integer, in string format.
func InsertCommaInt64(num int64) string {
	// func FormatUint(i uint64, base int) string
	str := strconv.FormatInt(num, 10)
	result := []string{}
	i := len(str) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(str, "") {
		if i == index {
			result = append(result, ",")
			i += 3
		}
		result = append(result, elem)
	}
	return strings.Join(result, "")
}

// InsertCommaFloat64 inserts comma in every three digit in integer part.
// It returns the new version of input float number, in string format.
func InsertCommaFloat64(num float64) string {
	// FormatFloat(num, 'f', 6, 64) with precision 6
	// for arbitrary precision, put -1
	str := strconv.FormatFloat(num, 'f', -1, 64)
	slice := strings.Split(str, ".")
	intpart := slice[0]
	floatpart := slice[1]
	result := []string{}
	i := len(intpart) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(intpart, "") {
		if i == index {
			result = append(result, ",")
			i += 3
		}
		result = append(result, elem)
	}

	intpart = strings.Join(result, "")
	return intpart + "." + floatpart
}
