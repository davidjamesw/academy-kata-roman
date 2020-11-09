package romanNumerals

var romanArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func convertRomanNumeralToArabic(roman string) int {
	return romanArabic[roman]
}
