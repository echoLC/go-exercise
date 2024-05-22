package exercise

import (
	"errors"
	"strings"
)

type RomanDec struct {
	decimal int
	roman   string
}

func divMod(numerator, divisor int) (int, int) {
	return numerator / divisor, numerator % divisor
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid input")
	} 

	romanDec := []RomanDec{
		{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
		{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
		{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
		{1000, "M"},
	}

	roman := ""
	n := 0
	for i := len(romanDec) - 1; 0 <= i; i-- {
		n, input = divMod(input, romanDec[i].decimal)
		roman += strings.Repeat(romanDec[i].roman, n)
	}
	return roman, nil
}
