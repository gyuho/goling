package spellcheck // import "github.com/gyuho/goling/spellcheck"

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Frequency counts the frequency of each word.
func Frequency(filename string) map[string]int {
	fmap := make(map[string]int)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	validID := regexp.MustCompile("[a-z]+")
	for _, word := range validID.FindAllString(strings.ToLower(string(text)), -1) {
		fmap[word]++
	}
	return fmap
}

/*
5. Run EditDistance 1
	Return all possible correction words c
		that has 1 edit distance from w
			that takes ONE edit to go from w to c
	Much more probable than the one with 2 edit distance
	but less probable than the one with 0 edit distance
	(Note that the results at this point
	 are not associated with the training data .
	 We just use the basic alphabetical combination
	 with const alphabet = "abcdefghijklmnopqrstuvwxyz" )
*/

// distance1 sends all possible corrections with edit distance 1 to the channel one.
func distance1(word string, one chan string) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	type pair struct {
		front, back string
	}

	splits := []pair{}
	for i := 0; i < len(word)+1; i++ {
		splits = append(splits, pair{word[:i], word[i:]})
	}

	for _, s := range splits {
		// deletion because we exclude s.back[0]
		if len(s.back) > 0 {
			one <- s.front + s.back[1:]
		}

		// transpose because we switch the elements in place
		if len(s.back) > 1 {
			one <- s.front + string(s.back[1]) + string(s.back[0]) + s.back[2:]
		}

		// replace because we insert one alphabet and exclude s.back[0]
		for _, elem := range alphabet {
			if len(s.back) > 0 {
				one <- s.front + string(elem) + s.back[1:]
			}
		}

		// insertion
		for _, elem := range alphabet {
			one <- s.front + string(elem) + s.back
		}
	}
}

/*
6. Run EditDistance 2
			Return all possible correction words c
					that has 2 edit distance from w
	To save time,
		Run another EditDistance 1
			in addition to the results of EditDistance 1
*/

// distances sends other possible corrections based on the results from distance1.
func distances(word string, other chan string) {
	one := make(chan string, 1024*1024)
	go func() {
		// Run distance1 first
		distance1(word, one)

		// And send empty string as breakpoint
		one <- ""
	}()
	// retrieve from distance1 results and break when it's done
	for e1 := range one {
		if e1 == "" {
			break
		}
		// run distance1 in addition to the results from the first distance1
		distance1(e1, other)
	}
}

// 7. To optimize, only extract the words that are in the training data

// knownWords returns the word with maximum frequencies.
func knownWords(str string, distFunc func(string, chan string), fmap map[string]int) string {
	words := make(chan string, 1024*1024)
	go func() {
		distFunc(str, words)
		words <- "" // send breakpoint
	}()
	maxFq := 0
	suggest := ""
	for wd := range words {
		if wd == "" {
			break
		}
		if freq, exist := fmap[wd]; exist && freq > maxFq {
			maxFq, suggest = freq, wd
		}
	}
	return suggest
}

// Suggest suggests the correct spelling based on the sample data.
func Suggest(word string, fmap map[string]int) string {
	// edit distance 0
	if _, exist := fmap[word]; exist {
		return word
	}

	// edit distance 1
	if sgst := knownWords(word, distance1, fmap); sgst != "" {
		return sgst
	}

	// edit distance 2
	if sgst := knownWords(word, distances, fmap); sgst != "" {
		return sgst
	}

	// edit distance 3, 4, and more ...
	return word
}
