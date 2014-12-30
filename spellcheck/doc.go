// Package spellcheck implements spell checker algorithms.
package spellcheck // import "github.com/gyuho/goling/spellcheck"

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
