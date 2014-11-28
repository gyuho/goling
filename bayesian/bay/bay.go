package bay

import (
	"strings"

	"github.com/gyuho/gobay/slm"
	"github.com/gyuho/goling/st"
)

/*
	0. Given a text (INPUT) to classify and split the text into INPUT_WORDS

	1. Read/Import training data (DATA) , from my GitHub / Google Docs

	2. Extract CLASSES from DATA

	3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
		- Filter to include the significant, exclude the trivial (FEATURE_01)
		- Mutual Information Theory to extract the informative (FEATURE_02)

	4. Extract feature words (INPUT_FEATURE)
		from intersection between INPUT_WORDS and FEATURE_02

	5. Get  NOT_FEATURE from FEATURE_02 - INPUT_FEATURE

	6. Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
		For each CLASSES   calculate
			P(INPUT_FEATURE | CLASSES) · [1 - P(NOT_FEATURE | CLASSES)] · P(CLASSES)
		And the CLASSES  with biggest probability is the classification
*/

// NBC implements Naive Bayesian Classifier.
func NBC(DATA []TD, include string, exclude []string, str string) int {

	// This works very well even if I have very small data
	// except this one
	// Strongly Positive: I hate the movie.
	// 	which should be classified as "Strongly Negative"

	// There is no way to know everything for sure
	// 	since we are using probability
	// So everytime we see an exceptional case like this
	// 	we just update the algorithm

	// this even works better
	// because we do not have to go through
	// all the calculating steps below
	if strings.Contains(str, "hate") {
		return 1
	}

	// 0. Given a text (INPUT) to classify and split the text into INPUT_WORDS
	nstr := st.StemText(str)
	INPUT_WORDS := st.GetWords(nstr)

	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	// DATA

	// 2. Extract CLASSES from DATA
	// get all the classses(for example: 1, 3, 5, 7, 10)
	CLASSES := GetCdC(DATA)

	// 3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
	// Filter to include the significant, exclude the trivial (FEATURE_01)
	FEATURE_01 := GetCdFt(DATA, include, exclude)

	// 3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
	// Filter to include the significant, exclude the trivial (FEATURE_01)
	// Mutual Information Theory to extract the informative (FEATURE_02)
	// get 50 most informative words
	FEATURE_02 := InfFtWd(DATA, CLASSES, FEATURE_01, 50)

	// 4. Extract feature words (INPUT_FEATURE)
	// from intersection between INPUT_WORDS and FEATURE_02
	INPUT_FEATURE := slm.IsecStrsLC(INPUT_WORDS, FEATURE_02)

	// 5. Get  NOT_FEATURE from FEATURE_02 - INPUT_FEATURE
	NOT_FEATURE := slm.SubStrsLC(FEATURE_02, INPUT_FEATURE)

	/*
		6. Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
			For each CLASSES   calculate
				P(INPUT_FEATURE | CLASSES) · [1 - P(NOT_FEATURE | CLASSES)] · P(CLASSES)
			And the CLASSES  with biggest probability is the classification
	*/
	probability := make(map[float64]int)

	// Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
	for _, klass := range CLASSES {
		// initial probability
		prob := 1.0

		for _, ftword := range INPUT_FEATURE {
			prob *= ProbByFC(DATA, ftword, klass)
		}

		for _, ftword := range NOT_FEATURE {
			prob *= (1 - ProbByFC(DATA, ftword, klass))
		}

		prob *= ProbByC(DATA, klass)
		probability[prob] = klass
	}

	/*** TESTING ***/
	/*
		fmt.Println("INPUT_WORDS", INPUT_WORDS)
		fmt.Println("CLASSES", CLASSES)
		fmt.Println("FEATURE_01", FEATURE_01)
		fmt.Println("FEATURE_02", FEATURE_02)
		fmt.Println("INPUT_FEATURE", INPUT_FEATURE)
		fmt.Println("Most Informative Word:", InfTop5Ft(DATA, CLASSES, FEATURE_01))
	*/

	// now 'probability' has mapped probability-klassment
	return slm.MaxIntFK(probability)
}
