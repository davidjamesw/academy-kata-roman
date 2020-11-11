package romannumerals

import (
	"errors"
	"fmt"
	"strings"
)

var romanArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var validReducers = map[string]string{
	"V": "I",
	"X": "I",
	"L": "X",
	"C": "X",
	"D": "C",
	"M": "C",
}

type romanArabicPair struct {
	roman  string
	arabic int
}

func convertRomanNumeralToArabic(romanNumerals string) (int, error) {
	splitRomans := strings.Split(romanNumerals, "")
	if error := validateRomanNumerals(splitRomans); error != nil {
		return -1, error
	}

	var total, arabicCurrentIterations int
	var previousPair romanArabicPair
	for i := len(splitRomans) - 1; i >= 0; i-- {
		currentPair := createPair(splitRomans[i])
		arabicCurrentIterations = numberOfIterations(currentPair.arabic, previousPair.arabic, arabicCurrentIterations)
		var nextPair romanArabicPair
		if i > 0 {
			nextPair = createPair(splitRomans[i-1])
		}
		if error := validateNumber(currentPair, previousPair, nextPair, arabicCurrentIterations); error != nil {
			return -1, error
		}
		total += numberToAdd(currentPair.arabic, previousPair.arabic)
		previousPair = currentPair
	}
	return total, nil
}

func createPair(roman string) romanArabicPair {
	return romanArabicPair{
		roman:  roman,
		arabic: romanArabic[roman],
	}
}

func numberOfIterations(arabicCurrent, arabicPrevious, arabicCurrentIterations int) int {
	if arabicCurrent != arabicPrevious {
		return 1
	}
	return arabicCurrentIterations + 1
}

func numberToAdd(arabicCurrent, arabicPrevious int) int {
	if arabicCurrent < arabicPrevious {
		return 0 - arabicCurrent
	}
	return arabicCurrent
}

func validateRomanNumerals(romanNumerals []string) error {
	for _, romanNumeral := range romanNumerals {
		if romanArabic[romanNumeral] == 0 {
			return errors.New("Roman Numeral contains at least one invalid character")
		}
	}
	return nil
}

func validateNumber(currentPair, previousPair, nextPair romanArabicPair, arabicCurrentIterations int) error {
	if arabicCurrentIterations == 4 {
		return errors.New("The same numeral can't be repeated more than three times in a row")
	}
	if arabicCurrentIterations == 2 && (currentPair.arabic == 5 || currentPair.arabic == 50 || currentPair.arabic == 500) {
		return errors.New("A five character can not be repeated")
	}
	if previousPair.arabic != 0 {
		validReducer := validReducers[previousPair.roman]
		if currentPair.arabic < previousPair.arabic && (validReducer != currentPair.roman) {
			errorMessage := fmt.Sprintf("%v can not be reduced by %v", previousPair.roman, currentPair.roman)
			return errors.New(errorMessage)
		}
		if nextPair.arabic != 0 {
			if currentPair.arabic < previousPair.arabic && nextPair.arabic <= currentPair.arabic {
				return errors.New("Reducing characters can not be repeated or reduced")
			}
			if (nextPair.arabic < currentPair.arabic) && (nextPair.arabic <= previousPair.arabic) {
				return errors.New("A reducing numeral can only be used when the numeral it is reducing hasn't already been increased by an equivalent or higher numeral. For example, IXI or VCX")
			}
		}
	}
	return nil
}
