package goling

/*
http://norvig.com/ngrams/ch14.pdf

Similar to Spell Check, we also use probability to find segmentation

1. Read the training data

2. By counting the frequency of each word
Calculate the probability of each word
The word with higher probability is more likely
to appear in the training data text

3. Do Smoothing over the words not in the corpus

4. Traverse all possible splits

5. Return the combination with highest probability
*/

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// 1. Read the training data

// OpenFile reads the training data.
func OpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		os.Exit(1)
	}
	return f
}

// CalculateWordProbability calculates probability
// : for each word, calculate how many times it occurs.
func CalculateWordProbability(reader io.Reader) map[string]float64 {
	// func NewScanner(r io.Reader) *Scanner
	// NewScanner returns a new Scanner to read from r.
	// The split function defaults to ScanLines.
	scanner := bufio.NewScanner(reader)

	// func (s *Scanner) Split(split SplitFunc)
	// Split sets the split function for the Scanner.
	// If called, it must be called before Scan.
	// The default split function is ScanLines.
	scanner.Split(bufio.ScanWords)

	// Map the words with probabilities.
	wordAndProb := make(map[string]float64)

	var length float64

	// func (*Scanner) Scan
	// Return false when the scan stops
	// either by reaching the end of the input or an error.
	for scanner.Scan() {

		// Remove all leading and trailing Unicode code points.
		// func Trim(s string, cutset string) string
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")

		_, exist := wordAndProb[word]
		/*
			2. By counting the frequency of each word
			Calculate the probability of each word
			The word with higher probability is more likely
			to appear in the training data text
		*/
		// if exists
		if exist {
			wordAndProb[word] += 1
		} else {
			wordAndProb[word] = 1
		}

		// keep increasing while reading(scanning)
		length++
	}

	// calculate the probability
	for str, floatvalue := range wordAndProb {
		wordAndProb[str] = floatvalue / length
	}
	return wordAndProb
}

// 3. Do Smoothing over the words not in the corpus

// CalculateWordNotInList calculates the probabilities
// of words that are not in the corpus.
func CalculateWordNotInList(word string, n int) float64 {
	length := float64(len(word))

	// prefer short strings by increasing the denominator
	return 10 / (float64(n) * math.Pow(10, length))
}

// ConstructWordProb constructs the word probability
// from the sample text file, and returns the probability.
func ConstructWordProb(reader io.Reader) func(string) float64 {
	wordprob_map := CalculateWordProbability(reader)

	return func(str string) float64 {
		score, exist := wordprob_map[str]
		if exist {
			return score
		}
		return CalculateWordNotInList(str, len(wordprob_map))
	}
}

type split struct {
	Head string
	Tail string
}

// PossibleSplits returns all possible splits.
func PossibleSplits(text string) []split {
	var result []split

	for i := range text {
		result = append(result, split{text[:i+1], text[i+1:]})
	}

	return result
}

// GetMostPlausible returns the most plausible combination
// , from the candidates.
func GetMostPlausible(words [][]string, getprobStr func(string) float64) []string {
	var max []string
	maximum_score := float64(-1)

	for _, candidate := range words {
		var totalscore float64 = 1
		for _, word := range candidate {
			totalscore *= getprobStr(word)
		}

		if maximum_score < totalscore {
			max = candidate
			maximum_score = totalscore
		}
	}

	return max
}

var BeenSeen = map[string][]string{}

// GetSegment returns the highest-scoring segmentation.
func GetSegment(text string, getprobStr func(string) float64) []string {

	if len(text) == 0 {
		return []string{}
	}

	result, exist := BeenSeen[text]
	if exist {
		return result
	}

	candidates := make([][]string, 0)

	// 4. Traverse all possible splits
	for _, thesplit := range PossibleSplits(text) {
		// append the string array
		candidates = append(candidates, append([]string{thesplit.Head}, GetSegment(thesplit.Tail, getprobStr)...))
	}

	// 5. Return the combination with highest probability
	max := GetMostPlausible(candidates, getprobStr)
	BeenSeen[text] = max
	return max
}
