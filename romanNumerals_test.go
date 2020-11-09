package romanNumerals

import (
	"testing"
)

func TestOne(t *testing.T) {
	testMatchesExpected("I", 1, t)
}

func TestFive(t *testing.T) {
	testMatchesExpected("V", 5, t)
}

func TestFifty(t *testing.T) {
	testMatchesExpected("L", 50, t)
}

func TestOneThousand(t *testing.T) {
	testMatchesExpected("M", 1000, t)
}

func testMatchesExpected(roman string, expected int, t *testing.T) {
	result := convertRomanNumeralToArabic(roman)
	if result != expected {
		t.Errorf("Was expecting %v but got %v", expected, result)
	}
}
