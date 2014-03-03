package st

import "strings"

// CountCharacter returns the total number of single characters.
func CountCharacter(str string) int {
	str = CleanUp(str)
	slice := strings.Split(str, "")
	return len(slice)
}

// CountWord returns the total number of words.
func CountWord(str string) int {
	slice := GetWords(str)
	return len(slice)
}

// CountSentence return the total number of sentences.
func CountSentence(str string) int {
	slice := SplitSent(str)
	return len(slice)
}
