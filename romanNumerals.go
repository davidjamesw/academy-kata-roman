package romanNumerals

import (
	"errors"
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

func convertRomanNumeralToArabic(romanNumerals string) (int, error) {
	splitRomans := strings.Split(romanNumerals, "")
	var total, arabicPrevious, arabicCurrentIterations int
	for i := len(splitRomans) - 1; i >= 0; i-- {
		arabicCurrent := romanArabic[splitRomans[i]]
		arabicCurrentIterations = numberOfIterations(arabicCurrent, arabicPrevious, arabicCurrentIterations)
		if error := validateNumber(arabicCurrent, arabicPrevious, arabicCurrentIterations); error != nil {
			return -1, error
		}
		total += nextNumber(arabicCurrent, arabicPrevious)
		arabicPrevious = arabicCurrent
	}
	return total, nil
}

func nextNumber(arabicCurrent, arabicPrevious int) int {
	if arabicCurrent < arabicPrevious {
		return 0 - arabicCurrent
	} else {
		return arabicCurrent
	}
}

func validateNumber(arabicCurrent, arabicPrevious, arabicCurrentIterations int) error {
	if arabicCurrentIterations == 4 {
		return errors.New("The same numeral can't be repeated more than three times in a row.")
	}
	if arabicCurrentIterations == 2 && (arabicCurrent == 5 || arabicCurrent == 50 || arabicCurrent == 500) {
		return errors.New("A five character can not be repeated")
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
