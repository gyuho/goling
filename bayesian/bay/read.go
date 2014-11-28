package bay

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gyuho/gobay/slm"
)

// GetStruct imports data from a csv file and construct the structure.
func GetStruct(filename string) []TD {
	output := CSV(filename)
	var stsli TD
	var result []TD

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		stsli.Class = slm.StrToInt(output[i][0])
		stsli.Weight = slm.StrToInt(output[i][0]) * 100
		stsli.Text = output[i][1]
		result = append(result, stsli)
	}
	return result
}

// GetInclFt imports "include" feature candidate range data from a csv file.
// Possibly big file, so use strings.Contains method should be faster.
func GetInclFt(filename string) string {
	output := CSV(filename)
	var icstr string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		icstr += "," + output[i][0]
	}
	return icstr
}

// GetExcFt imports "exclude" feature candidate range data from a csv file.
// Relatively small amount of data.
// Just to be used with linear search.
func GetExcFt(filename string) []string {
	output := CSV(filename)
	var exclude_arr []string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		exclude_arr = append(exclude_arr, output[i][0])
	}

	return exclude_arr
}

// CSV reads data from a csv file.
// [][]string:
// the first [] is row, the second [] is column.
// len(output) would be the number of total rows.
// Use the following line to traverse
// by all rows and only the first column.
// for i := 0; i < len(output); i++
// 	output[i][0]
func CSV(filename string) [][]string {
	// func Open(name string) (file *File, err error)
	file, err := os.Open(filename)
	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println("Error:", err)
		return nil
	}

	// func (f *File) Close() error
	defer file.Close()

	// func NewReader(r io.Reader) *Reader
	// it can read csv or txt file
	reader := csv.NewReader(file)

	reader.TrailingComma = true
	reader.TrimLeadingSpace = true
	// reader.LazyQuotes = true

	for {
		// func (r *Reader) ReadAll() (records [][]string, err error)
		data, read_err := reader.ReadAll()

		if read_err == io.EOF {
			break
		} else if read_err != nil {
			// fmt.Println(err.Error())
			fmt.Println("Error:", read_err)
			return nil
		}
		return data
	}
	return nil
}
