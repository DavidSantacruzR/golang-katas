package fibonacci

import (
	"reflect"
	"testing"
	"time"
)

func TestGetFibonacciSequenceUpTo0(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(0)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo1(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(1)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo2(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(2)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo3(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(3)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo4(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(4)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo5(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(5)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo6(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(6)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo7(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(7)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo8(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(8)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo9(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(9)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo10(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(10)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}

func TestGetFibonacciSequenceUpTo40(t *testing.T) {
	start := time.Now()
	result := GetFibonacciSequence(40)
	elapsed := time.Since(start)
	t.Logf("Execution time: %s", elapsed)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFibonacciSequence(5): got %v; want %v", result, expected)
	}
}
