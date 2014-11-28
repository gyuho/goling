package bay

import "fmt"

// Print prints out the outcome.
func Print(DATA []TD, include string, exclude []string, str string) {
	result := NBC(DATA, include, exclude, str)
	switch result {
	case 1:
		fmt.Println("Strongly Negative:", str)
	case 2:
		fmt.Println("Very Negative:", str)
	case 3:
		fmt.Println("Negative:", str)
	case 4:
		fmt.Println("Little Negative:", str)
	case 5:
		fmt.Println("Neutral:", str)
	case 6:
		fmt.Println("Little Positive:", str)
	case 7:
		fmt.Println("Positive:", str)
	case 8:
		fmt.Println("More Postivie:", str)
	case 9:
		fmt.Println("Very Positive:", str)
	case 10:
		fmt.Println("Strongly Positive:", str)
	default:
		fmt.Println("Failed to detect:", str)
	}
}
