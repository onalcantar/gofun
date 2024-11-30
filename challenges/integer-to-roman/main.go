package main

import (
	"strings"
)

var romansList = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	if num <= 0 {
		return ""
	}

	var romanStr strings.Builder

    for _, roman := range romansList {
		if num >= roman.Value {
			timesToAdd := num / roman.Value
			romanStr.WriteString(strings.Repeat(roman.Symbol, timesToAdd))
			num -= roman.Value * timesToAdd
		}
	}

	return romanStr.String()
}

