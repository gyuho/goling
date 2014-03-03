package seg

import (
	"fmt"
	"os"
)

// OpenFile reads the training data.
func OpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		os.Exit(1)
	}
	return f
}
