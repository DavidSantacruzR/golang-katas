package palindrome

import (
	"strings"
)

func ReverseString(input string) string {
	result := make([]string, len(input))
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = string(input[j]), string(input[i])
	}
	return strings.Join(result, "")
}

func IsPalindrome(word string) bool {
	modulo := len(word) % 2
	midPoint := len(word) / 2
	if modulo != 0 {
		return strings.ToLower(word[:midPoint]) == strings.ToLower(ReverseString(word[midPoint+1:]))
	} else {
		return strings.ToLower(word[:midPoint]) == strings.ToLower(ReverseString(word[midPoint:]))
	}
}
