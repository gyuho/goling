package segment

import (
	"bufio"
	"io"
	"math"
	"strings"
)

// probability calculates frequency probability.
// For each word, divide frequency by the number of words.
func probability(reader io.Reader) map[string]float64 {
	scanner := bufio.NewScanner(reader)
	//
	// This must be called before Scan.
	// The default split function is bufio.ScanLines.
	scanner.Split(bufio.ScanWords)
	//
	pmap := make(map[string]float64)
	//
	var length float64
	for scanner.Scan() {
		// Remove all leading and trailing Unicode code points.
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")
		if _, exist := pmap[word]; exist {
			pmap[word]++
		} else {
			pmap[word] = 1
		}
		// keep increasing while reading(scanning)
		length++
	}
	for k, v := range pmap {
		pmap[k] = v / length
	}
	return pmap
}

// Probability returns the word probability,
// with smoothing.
func Probability(reader io.Reader) func(string) float64 {
	pmap := probability(reader)
	return func(word string) float64 {
		if score, ok := pmap[word]; ok {
			return score
		}
		// if the word has never showed up, smooth.
		return 10 / (float64(len(pmap)) * math.Pow(10, float64(len(word))))
	}
}
