package segment

import (
	"bufio"
	"io"
	"math"
	"strings"
)

// http://norvig.com/ngrams/ch14.pdf
// Similar to Spell Check, we also use probability to find segmentation.
//
// 1. Read the training data.
//
// 2. By counting the frequency of each word, calculate the probability
// of each word. The word with higher probability is more likely to appear
// in the training data text.
//
// 3. Do Smoothing over the words not in the corpus.
//
// 4. Traverse all possible splits.
//
// 5. Return the combination with highest probability.

// frequency calculates frequency probability.
// For each word, divide frequency by the number of words.
func frequency(reader io.Reader) map[string]float64 {
	scanner := bufio.NewScanner(reader)

	// func (s *Scanner) Split(split SplitFunc)
	// Split sets the split function for the Scanner.
	// If called, it must be called before Scan.
	// The default split function is ScanLines.
	scanner.Split(bufio.ScanWords)

	// Map the words with probabilities.
	rmap := make(map[string]float64)

	var length float64

	// func (*Scanner) Scan
	// Return false when the scan stops
	// either by reaching the end of the input or an error.
	for scanner.Scan() {
		// Remove all leading and trailing Unicode code points.
		// func Trim(s string, cutset string) string
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")
		if _, exist := rmap[word]; exist {
			rmap[word]++
		} else {
			rmap[word] = 1
		}
		// keep increasing while reading(scanning)
		length++
	}
	for k, v := range rmap {
		rmap[k] = v / length
	}
	return rmap
}

// Frequency constructs the word probability.
func Frequency(reader io.Reader) func(string) float64 {
	fmap := frequency(reader)
	return func(word string) float64 {
		if score, ok := fmap[word]; ok {
			return score
		}
		// if the word has never showed up, smoothing
		return 10 / (float64(len(fmap)) * math.Pow(10, float64(len(word))))
	}
}

func mostPlausible(words [][]string, probFunc func(string) float64) []string {
	max := []string{}
	min := float64(-1)
	for _, row := range words {
		score := 1.0
		for _, elem := range row {
			score *= probFunc(elem)
		}
		if min < score {
			max = row
			min = score
		}
	}
	return max
}

type split struct {
	Head string
	Tail string
}

func getsplit(str string) []split {
	splits := []split{}
	for i := range str {
		splits = append(splits, split{str[:i+1], str[i+1:]})
	}
	return splits
}

// beenSeen stores segmentations that have been seen.
var beenSeen = map[string][]string{}

// Get returns the highest-scoring segmentation.
func Get(text string, myfunc func(string) float64) []string {
	if len(text) == 0 {
		return []string{}
	}
	if result, ok := beenSeen[text]; ok {
		return result
	}

	candidates := [][]string{}
	for _, thesplit := range getsplit(text) {
		candidates = append(
			candidates,
			append(
				[]string{thesplit.Head},
				Get(thesplit.Tail, myfunc)...,
			),
		)
	}

	max := mostPlausible(candidates, myfunc)
	beenSeen[text] = max
	return max
}
