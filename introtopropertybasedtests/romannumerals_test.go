package introtopropertybasedtests

import (
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Arabic      uint16
	Roman       string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"3 gets converted to III", 3, "III"},
	{"4 gets converted to IV", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"6 gets converted to VI", 6, "VI"},
	{"7 gets converted to VII", 7, "VII"},
	{"8 gets converted to VIII", 8, "VIII"},
	{"9 gets converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{"50 gets converted to L", 50, "L"},
	{"90 gets converted to XC", 90, "XC"},
	{"99 gets converted to XCIX", 99, "XCIX"},
	{"100 gets converted to C", 100, "C"},
	{"499 gets converted to CDXCIX", 499, "CDXCIX"},
	{"400 gets converted to CD", 400, "CD"},
	{"500 gets converted to D", 500, "D"},
	{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
	{"999 gets converted to CMXCIX", 999, "CMXCIX"},
	{"1000 gets converted to M", 1000, "M"},
	{"1006 gets converted to MVI", 1006, "MVI"},
	{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
	{"2014 gets converted to MMXIV", 2014, "MMXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, but wanted %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, but wanted %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func TestPropertiesOfConsecutiveSymbols(t *testing.T) {
	disallowedConsecutiveSymbols := []string{
		"IIII",
		"VVVV",
		"XXXX",
		"MMMM",
		"CCCC",
		"DDDD",
		"LLLL",
	}

	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(arabic)
		t.Log("testing", roman, arabic)

		for _, symbol := range disallowedConsecutiveSymbols {
			if strings.Contains(roman, symbol) {
				return false
			}
		}
		return true
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
