package romannumerals

import (
	"fmt"
	"strings"
	"testing"
)

func TestIReturns1(t *testing.T) {
	testMatchesExpected("I", 1, t)
}

func TestIIReturns2(t *testing.T) {
	testMatchesExpected("II", 2, t)
}

func TestIIIReturns3(t *testing.T) {
	testMatchesExpected("III", 3, t)
}

func TestIVReturns4(t *testing.T) {
	testMatchesExpected("IV", 4, t)
}

func TestVReturns5(t *testing.T) {
	testMatchesExpected("V", 5, t)
}

func TestIVReturns6(t *testing.T) {
	testMatchesExpected("VI", 6, t)
}

func TestIXReturns9(t *testing.T) {
	testMatchesExpected("IX", 9, t)
}

func TestLReturns50(t *testing.T) {
	testMatchesExpected("L", 50, t)
}

func TestCCReturns200(t *testing.T) {
	testMatchesExpected("CC", 200, t)
}

func TestCDReturns400(t *testing.T) {
	testMatchesExpected("CD", 400, t)
}

func TestDReturns500(t *testing.T) {
	testMatchesExpected("D", 500, t)
}
func TestCMReturns900(t *testing.T) {
	testMatchesExpected("CM", 900, t)
}
func TestMReturns1000(t *testing.T) {
	testMatchesExpected("M", 1000, t)
}

func TestMMVIReturns2006(t *testing.T) {
	testMatchesExpected("MMVI", 2006, t)
}

func TestTMMMCMXCIXReturns3999(t *testing.T) {
	testMatchesExpected("MMMCMXCIX", 3999, t)
}

func testMatchesExpected(roman string, expected int, t *testing.T) {
	result, _ := convertRomanNumeralToArabic(roman)
	if result != expected {
		t.Errorf("Was expecting %v but got %v", expected, result)
	}
}

const tooManyTimesError = "The same numeral can't be repeated more than three times in a row"
const repeatedFivesError = "A five character can not be repeated"
const repeatedReducerError = "Reducing characters can not be repeated or reduced"

func TestTooManyConsecutivesGivesError(t *testing.T) {
	testErrorThrown("IIII", tooManyTimesError, t)
}

func TestTooManyConsecutivesInMiddleOfNumberGivesError(t *testing.T) {
	testErrorThrown("MCCCCX", tooManyTimesError, t)
}

func TestTooManyConsecutiveFivesGivesError(t *testing.T) {
	testErrorThrown("VV", repeatedFivesError, t)
}

func TestTooManyConsecutiveFiftiesGivesError(t *testing.T) {
	testErrorThrown("LL", repeatedFivesError, t)
}

func TestTooManyReducersTenGivesError(t *testing.T) {
	testErrorThrown("IIX", repeatedReducerError, t)
}

func TestTooManyReducersXCMGivesError(t *testing.T) {
	testErrorThrown("XCM", repeatedReducerError, t)
}

func testErrorThrown(roman, message string, t *testing.T) {
	res, err := convertRomanNumeralToArabic(roman)
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
		return
	}
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}

func TestInvalidReducerXMGivesError(t *testing.T) {
	testInvalidReducer("XM", t)
}

func TestInvalidReducerCVXGivesError(t *testing.T) {
	testInvalidReducer("CVX", t)
}

func TestInvalidReducerILGivesError(t *testing.T) {
	testInvalidReducer("IL", t)
}

func TestInvalidReducerIMGivesError(t *testing.T) {
	testInvalidReducer("IM", t)
}

func testInvalidReducer(roman string, t *testing.T) {
	res, err := convertRomanNumeralToArabic(roman)
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
		return
	}
	splitRoman := strings.Split(roman, "")
	roman1 := splitRoman[len(splitRoman)-1]
	roman2 := splitRoman[len(splitRoman)-2]
	message := fmt.Sprintf("%v can not be reduced by %v", roman1, roman2)
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}

func TestInvalidCharactersGivesError(t *testing.T) {
	res, err := convertRomanNumeralToArabic("THISISINVALID")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
		return
	}
	message := "Roman Numeral contains at least one invalid character"
	if err.Error() != message {
		t.Errorf("Error message was %v but expected %v", err.Error(), message)
	}
}

var invalidReducingCharacterMessage = "A reducing numeral can only be used when the numeral it is reducing hasn't already been increased by an equivalent or higher numeral. For example, IXI or VCX"

func TestReducingAnIncreasedValueGivesErrorIXI(t *testing.T) {
	res, err := convertRomanNumeralToArabic("IXI")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
		return
	}
	if err.Error() != invalidReducingCharacterMessage {
		t.Errorf("Error message was %v but expected %v", err.Error(), invalidReducingCharacterMessage)
	}
}

func TestReducingAnIncreasedValueGivesErrorCMCIX(t *testing.T) {
	res, err := convertRomanNumeralToArabic("CMCIX")
	if err == nil {
		t.Errorf("Expected an error but got %v", res)
		return
	}
	if err.Error() != invalidReducingCharacterMessage {
		t.Errorf("Error message was %v but expected %v", err.Error(), invalidReducingCharacterMessage)
	}
}
