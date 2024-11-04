package cipher

import "testing"

func TestCeasarEncodingShiftOdd(t *testing.T) {
	result := CeasarEncoding("david", 3)
	expected := "gdylg"
	if result != expected {
		t.Errorf("CeasarEncoding: got %v; want %v", result, expected)
	}
}

func TestCeasarEncodingShiftEven(t *testing.T) {
	result := CeasarEncoding("david", 4)
	expected := "hezmh"
	if result != expected {
		t.Errorf("CeasarEncoding: got %v; want %v", result, expected)
	}
}

func TestCeasarDecoding(t *testing.T) {
	result := CeasarDecoding("gdylg", 3)
	expected := "david"
	if result != expected {
		t.Errorf("CeasarDecoding: got %v; want %v", result, expected)
	}
}
