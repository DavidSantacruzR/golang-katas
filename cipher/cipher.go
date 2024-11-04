package cipher

import (
	"strings"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func CeasarEncoding(message string, shift int) string {
	result := make([]string, len(message))
	alphabetLength := len(Alphabet)
	for i := 0; i < len(message); i++ {
		lookupIndex := (strings.Index(Alphabet, string(message[i])) + shift) % alphabetLength
		result[i] = string(Alphabet[lookupIndex])
	}
	return strings.Join(result, "")
}

func CeasarDecoding(message string, shift int) string {
	result := make([]string, len(message))
	alphabetLength := len(Alphabet)
	for i := 0; i < len(message); i++ {
		lookupIndex := (strings.Index(Alphabet, string(message[i])) - shift) % alphabetLength
		result[i] = string(Alphabet[lookupIndex])
	}
	return strings.Join(result, "")
}
