package morestrings

import "strings"

// Get intersection between two strings
func Intersection(str1, str2 string) string {
	result := ""

	for _, letter := range str2 {
		if strings.Contains(str1, string(letter)) {
			result += string(letter)
		}
	}

	return result
}

// Remove duplicate character in a string
// Return string with distinct characters
func RemoveDups(str string) string {
	result := string(str[0])

	for _, letter := range str {
		if !strings.Contains(result, string(letter)) {
			result += string(letter)
		}
	}

	return result
}
