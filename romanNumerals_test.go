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

func TestThree(t *testing.T) {
	testMatchesExpected("III", 3, t)
}

func TestFour(t *testing.T) {
	testMatchesExpected("IV", 4, t)
}

func TestFive(t *testing.T) {
	testMatchesExpected("V", 5, t)
}

func TestSix(t *testing.T) {
	testMatchesExpected("VI", 6, t)
}

func TestNine(t *testing.T) {
	testMatchesExpected("IX", 9, t)
}

func TestFifty(t *testing.T) {
	testMatchesExpected("L", 50, t)
}

func TestTwoHundred(t *testing.T) {
	testMatchesExpected("CC", 200, t)
}

func TestFourHundred(t *testing.T) {
	testMatchesExpected("CD", 400, t)
}

func TestFiveHundred(t *testing.T) {
	testMatchesExpected("D", 500, t)
}
func TestNineHundred(t *testing.T) {
	testMatchesExpected("CM", 900, t)
}
func TestOneThousand(t *testing.T) {
	testMatchesExpected("M", 1000, t)
}

func TestTwoThousandAndSix(t *testing.T) {
	testMatchesExpected("MMVI", 2006, t)
}

func TestTooManyConsecutives(t *testing.T) {
	res, err := convertRomanNumeralToArabic("IIII")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
}

func TestTooManyConsecutivesInMiddleOfNumber(t *testing.T) {
	res, err := convertRomanNumeralToArabic("MCCCCX")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
}

func TestTooManyConsecutiveFives(t *testing.T) {
	res, err := convertRomanNumeralToArabic("VV")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
}

func TestTooManyConsecutiveFifties(t *testing.T) {
	res, err := convertRomanNumeralToArabic("LL")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
}

func TestThreeThousandNineHundredAndNinetyNine(t *testing.T) {
	testMatchesExpected("MMMCMXCIX", 3999, t)
}

func testMatchesExpected(roman string, expected int, t *testing.T) {
	result, _ := convertRomanNumeralToArabic(roman)
	if result != expected {
		t.Errorf("Was expecting %v but got %v", expected, result)
	}
}
