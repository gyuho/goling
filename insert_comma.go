package goling

import (
	"strconv"
	"strings"
)

// InsertCommaInt inserts comma in every three digit.
// It returns the new version of input integer, in string format.
func InsertCommaInt(num int64) string {
	// func FormatUint(i uint64, base int) string
	str := strconv.FormatInt(num, 10)
	var numslice []string
	i := len(str) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(str, "") {
		if i == index {
			numslice = append(numslice, ",")
			i += 3
		}
		numslice = append(numslice, elem)
	}
	return strings.Join(numslice, "")
}

// InsertCommaFloat inserts comma in every three digit in integer part.
// It returns the new version of input float number, in string format.
func InsertCommaFloat(num float64) string {
	// FormatFloat(num, 'f', 6, 64) with precision 6
	// for arbitrary precision, put -1
	str := strconv.FormatFloat(num, 'f', -1, 64)

	slice := strings.Split(str, ".")

	intpart := slice[0]
	floatpart := slice[1]

	var intresult []string
	i := len(intpart) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(intpart, "") {
		if i == index {
			intresult = append(intresult, ",")
			i += 3
		}
		intresult = append(intresult, elem)
	}

	intpart = strings.Join(intresult, "")
	return intpart + "." + floatpart
}
