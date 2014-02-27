package goling

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// TRAIN trains sample data so that we can retrieve
// the likely correct spelling from the sample data.
func TRAIN(filename string) map[string]int {
	NWORDS := make(map[string]int)
	validID := regexp.MustCompile("[a-z]+")
	if text, err := ioutil.ReadFile(filename); err == nil {
		// convert to lower case
		for _, word := range validID.FindAllString(strings.ToLower(string(text)), -1) {
			// increase its frequency by 1
			// everytime we see the word
			NWORDS[word]++
		}
	} else {
		panic("Try again. Failed to load training data!")
	}
	return NWORDS
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

// EditDistance1 sends all possible corrections
// with edit distance 1 to the channel ch.
func EditDistance1(word string, ch chan string) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	type Pair struct {
		front, back string
	}

	var splits []Pair
	for i := 0; i < len(word)+1; i++ {
		splits = append(splits, Pair{word[:i], word[i:]})
	}

	for _, s := range splits {
		// deletion because we exclude s.back[0]
		if len(s.back) > 0 {
			ch <- s.front + s.back[1:]
		}

		// transpose because we switch the elements in place
		if len(s.back) > 1 {
			ch <- s.front + string(s.back[1]) + string(s.back[0]) + s.back[2:]
		}

		// replace because we insert one alphabet and exclude s.back[0]
		for _, elem := range alphabet {
			if len(s.back) > 0 {
				ch <- s.front + string(elem) + s.back[1:]
			}
		}

		// insertion
		for _, elem := range alphabet {
			ch <- s.front + string(elem) + s.back
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

// EditDistance2 sends other possible corrections
// based on the results from EditDistance1.
func EditDistance2(word string, ch chan string) {
	ch_from_EditDistance1 := make(chan string, 1024*1024)
	go func() {
		// Run EditDistance1 first
		EditDistance1(word, ch_from_EditDistance1)

		// And send empty string as breakpoint
		ch_from_EditDistance1 <- ""
	}()

	// retrieve from EditDistance1 results
	// and break when it's done
	for e1 := range ch_from_EditDistance1 {
		if e1 == "" {
			break
		}

		// run EditDistance1 in addition to the results
		// from the first EditDistance1
		EditDistance1(e1, ch)
	}
}

// 7. To optimize, only extract the words that are in the training data

// KnownWord returns the word with maximum frequencies.
// If we want to return more than one candidates
// we can just use slice to store the results.
func KnownWord(str string, edits func(string, chan string), model map[string]int) string {
	ch := make(chan string, 1024*1024)
	go func() {
		edits(str, ch)
		// send breakpoint
		ch <- ""
	}()
	maxFq := 0
	correction := ""
	for c := range ch {
		if c == "" {
			break
		}
		if freq, exist := model[c]; exist && freq > maxFq {
			maxFq, correction = freq, c
		}
	}
	return correction
}

// 8. c with less edit distance is much more probable

// SpellCorrect corrects the spelling based on the sample data.
func SpellCorrect(word string, model map[string]int) string {
	// edit distance 0
	if _, exist := model[word]; exist {
		return word
	}

	// edit distance 1
	if correction := KnownWord(word, EditDistance1, model); correction != "" {
		return correction
	}

	// edit distance 2
	if correction := KnownWord(word, EditDistance2, model); correction != "" {
		return correction
	}

	// edit distance 3, 4, and more ...
	return word
}

/*
norvig.com/spell-correct.html
cxwangyi.wordpress.com

1. Train spelling pattern with text data
2. Choose the most likely spelling correction for the word

Since there is no way to know for sure,
we want to use probability

Suppose we want to find
the correction, or correct spelling, c
out of all possible corrections
that maximizes the probability of correction c
given the original, wrong-spelled, word w

In probability
argmax _x f(x) means
 the set of x values
  for which f(x) attains its largest value

Then we need to find
argmax _c  P(c | w)

when P(c) is the probability
 that a proposed correction c stands on itw own
	That is, P(c) is the probability
	 that c appears in the training data
	 		P("the") should be higher than P("ewrsdfsd")

P(w) is the probability
 that the wrong-spelled word w appears in text


By Bayes' Theorem
P(A | B) = P(B|A)·P(A)/P(B)
P(c | w) = P(w|c)·P(c)/P(w)
 but P(w) is always the same

Therefore we only need to find P(w|c)·P(c)

argmax _c P(w|c)·P(c)
means that we enumerate all feasible values of c
and then choose the one with the highest probability

The reason why we use this is
to consider separate two factors
when calculating the probabilities

argmax _c P(w|c)·P(c)

P(c) = probability that c appears in the training data(TEXT)

1. TEXT: Download any data from gutenberg.org

2. NWORDS: Extract all individual words from the TEXT

3. TRAIN: Count the frequencies of each word (map data structure)

4. Do smoothing over the words that are not in the training data
		assigning the smallest probabilities
			as if we had seen them only once

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

6. Run EditDistance 2
			Return all possible correction words c
					that has 2 edit distance from w
	To save time,
		Run another EditDistance 1
			in addition to the results of EditDistance 1
	(Note that the results at this point
	 are not associated with the training data .
	 We just use the basic alphabetical combination
	 with const alphabet = "abcdefghijklmnopqrstuvwxyz" )

7. To optimize, only extract the words that are in the training data
		For example, we have 114,324 words of EditDistance2 from "something"
		but we can narrow it down to 4 ~ 5 words
			by limiting the candidates to be found
			 only in the training data.

8. To optimize even more,
		consider that one vowel edit
		 is always more probable than TWO consonant edits
	The c with less edit distance is much more probable
	And c with 0 edit distance means there is no need to edit.
	Then we just return this right away.
*/
