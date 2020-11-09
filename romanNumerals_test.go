package romanNumerals

import (
	"testing"
)

func TestOne(t *testing.T) {
	testMatchesExpected("I", 1, t)
}

func TestTwo(t *testing.T) {
	testMatchesExpected("II", 2, t)
}

func TestFive(t *testing.T) {
	testMatchesExpected("V", 5, t)
}

func TestSix(t *testing.T) {
	testMatchesExpected("VI", 6, t)
}

func TestFifty(t *testing.T) {
	testMatchesExpected("L", 50, t)
}

func TestOneThousand(t *testing.T) {
	testMatchesExpected("M", 1000, t)
}

func TestTwoThousandAndSix(t *testing.T) {
	testMatchesExpected("MMVI", 2006, t)
}

func testMatchesExpected(roman string, expected int, t *testing.T) {
	result := convertRomanNumeralToArabic(roman)
	if result != expected {
		t.Errorf("Was expecting %v but got %v", expected, result)
	}
}
