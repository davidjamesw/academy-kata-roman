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

func TestThreeThousandNineHundredAndNinetyNine(t *testing.T) {
	testMatchesExpected("MMMCMXCIX", 3999, t)
}

const tooManyTimesError = "The same numeral can't be repeated more than three times in a row."
const repeatedFivesError = "A five character can not be repeated"
const repeatedReducerError = "Reducing characters can not be repeated or reduced"

func TestTooManyConsecutives(t *testing.T) {
	testErrorThrown("IIII", tooManyTimesError, t)
}

func TestTooManyConsecutivesInMiddleOfNumber(t *testing.T) {
	testErrorThrown("MCCCCX", tooManyTimesError, t)
}

func TestTooManyConsecutiveFives(t *testing.T) {
	testErrorThrown("VV", repeatedFivesError, t)
}

func TestTooManyConsecutiveFifties(t *testing.T) {
	testErrorThrown("LL", repeatedFivesError, t)
}

func TestTooManyReducersTen(t *testing.T) {
	testErrorThrown("IIX", repeatedReducerError, t)
}

func TestTooManyReducersXCM(t *testing.T) {
	testErrorThrown("XCM", repeatedReducerError, t)
}

func testMatchesExpected(roman string, expected int, t *testing.T) {
	result, _ := convertRomanNumeralToArabic(roman)
	if result != expected {
		t.Errorf("Was expecting %v but got %v", expected, result)
	}
}

func testErrorThrown(roman, message string, t *testing.T) {
	res, err := convertRomanNumeralToArabic(roman)
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}

func TestInvalidReducerXM(t *testing.T) {
	res, err := convertRomanNumeralToArabic("XM")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
	message := "M can not be reduced by X"
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}

func TestInvalidReducerCVX(t *testing.T) {
	res, err := convertRomanNumeralToArabic("CVX")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
	}
	message := "X can not be reduced by V"
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}
