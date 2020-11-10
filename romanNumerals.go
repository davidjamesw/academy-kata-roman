package romanNumerals

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
	var total, arabicCurrentIterations int
	var previousReduced bool
	var previousPair romanArabicPair
	for i := len(splitRomans) - 1; i >= 0; i-- {
		currentPair := romanArabicPair{
			roman:  splitRomans[i],
			arabic: romanArabic[splitRomans[i]],
		}
		arabicCurrentIterations = numberOfIterations(currentPair.arabic, previousPair.arabic, arabicCurrentIterations)
		if error := validateNumber(currentPair, previousPair, arabicCurrentIterations, previousReduced); error != nil {
			return -1, error
		}
		var arabicNumberToAdd int
		arabicNumberToAdd, previousReduced = numberToAdd(currentPair.arabic, previousPair.arabic)
		total += arabicNumberToAdd
		previousPair = currentPair
	}
	return total, nil
}

func numberToAdd(arabicCurrent, arabicPrevious int) (int, bool) {
	if arabicCurrent < arabicPrevious {
		return 0 - arabicCurrent, true
	} else {
		return arabicCurrent, false
	}
}

func validateNumber(currentPair, previousPair romanArabicPair, arabicCurrentIterations int, previousReduced bool) error {
	if arabicCurrentIterations == 4 {
		return errors.New("The same numeral can't be repeated more than three times in a row.")
	}
	if arabicCurrentIterations == 2 && (currentPair.arabic == 5 || currentPair.arabic == 50 || currentPair.arabic == 500) {
		return errors.New("A five character can not be repeated")
	}
	validReducer := validReducers[previousPair.roman]
	if currentPair.arabic < previousPair.arabic && (validReducer != currentPair.roman) {
		errorMessage := fmt.Sprintf("%v can not be reduced by %v", previousPair.roman, currentPair.roman)
		return errors.New(errorMessage)
	}
	if currentPair.arabic <= previousPair.arabic && previousReduced {
		return errors.New("Reducing characters can not be repeated or reduced")
	}
	return nil
}

func numberOfIterations(arabicCurrent, arabicPrevious, arabicCurrentIterations int) int {
	if arabicCurrent != arabicPrevious {
		return 1
	} else {
		return arabicCurrentIterations + 1
	}
}
