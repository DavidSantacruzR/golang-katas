package main

import (
	"fmt"
	"golang/cipher"
	"golang/fibonacci"
	"golang/palindrome"
)

func main() {
	fmt.Println(fibonacci.GetFibonacciSequence(4))
	fmt.Println(palindrome.IsPalindrome("level"))
	fmt.Println(cipher.CeasarEncoding("david", 4))
}
