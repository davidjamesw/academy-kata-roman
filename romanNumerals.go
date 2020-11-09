package romanNumerals

import "strings"

var romanArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func convertRomanNumeralToArabic(romanNumerals string) int {
	splitRomans := strings.Split(romanNumerals, "")
	var total = 0

	for i := len(splitRomans) - 1; i >= 0; i-- {
		total += romanArabic[splitRomans[i]]
	}

	return total
}
