package utils

import "regexp"

// AllTabIntoSingleSpace converts all tab characters into single whitespace character.
func AllTabIntoSingleSpace(str string) string {
	// to take any tab chracters: single tab, double tabs, ...
	validID := regexp.MustCompile(`\t{1,}`)
	return validID.ReplaceAllString(str, " ")
}

// AllSpaceIntoSingleTab converts all whitespace characters into single tab character.
func AllSpaceIntoSingleTab(str string) string {
	// to take any whitespace characters: single whitespace, doulbe _, ...
	validID := regexp.MustCompile(`\s{1,}`)
	return validID.ReplaceAllString(str, "	")
}

// TabToSpace converts all tab characters into whitespace characters.
func TabToSpace(str string) string {
	validID := regexp.MustCompile(`\t`)
	return validID.ReplaceAllString(str, " ")
}

// SpaceToTab converts all whitespace characters into tab characters.
func SpaceToTab(str string) string {
	validID := regexp.MustCompile(`\s`)
	return validID.ReplaceAllString(str, "	")
}

// DeleteSpace removes all whitespace characters.
func DeleteSpace(str string) string {
	validID := regexp.MustCompile(`\s{1,}`)
	return validID.ReplaceAllString(str, "")
}

// DeleteTab removes all tab characters.
func DeleteTab(str string) string {
	validID := regexp.MustCompile(`\t{1,}`)
	return validID.ReplaceAllString(str, "")
}
