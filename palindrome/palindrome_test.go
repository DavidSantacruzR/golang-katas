package palindrome

import (
	"testing"
	"time"
)

func TestPalindromeCheckerOdd(t *testing.T) {
	start := time.Now()
	result := IsPalindrome("level")
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	if result != true {
		t.Error("Expected true, got ", result)
	}
}

func TestPalindromeCheckerEven(t *testing.T) {
	start := time.Now()
	result := IsPalindrome("noon")
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	if result != true {
		t.Error("Expected true, got ", result)
	}
}

func TestPalindromeCheckerFalse(t *testing.T) {
	start := time.Now()
	result := IsPalindrome("david")
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	if result != false {
		t.Error("Expected false, got ", result)
	}
}

func TestPalindromeCheckerEmptyString(t *testing.T) {
	start := time.Now()
	result := IsPalindrome("")
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	if result != true {
		t.Error("Expected false, got ", result)
	}
}

func TestPalindromeCheckerCamelCaseStrings(t *testing.T) {
	start := time.Now()
	result := IsPalindrome("Civic")
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	if result != true {
		t.Error("Expected false, got ", result)
	}
}
