package seg

import (
	"log"
	"os"
)

// OpenFile reads the training data.
func OpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	return f
}
