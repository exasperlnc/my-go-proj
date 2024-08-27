package romancata

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Num         int
		Want        string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 converts to III", 3, "III"},
		{"4 converts to IV", 4, "IV"},
		{"5 converts to V", 5, "V"},
		{"6 converts to VI", 6, "VI"},
		{"7 converts to VII", 7, "VII"},
		{"8 converts to VIII", 8, "VIII"},
		{"9 converts to IX", 9, "IX"},
		{"10 converts to X", 10, "X"},
		{"14 converts to XIV", 14, "XIV"},
		{"18 converts to XVIII", 18, "XVIII"},
		{"20 converts to XX", 20, "XX"},
		{"39 converts to XXXIX", 39, "XXXIX"},
		{"40 converts to XL", 40, "XL"},
		{"47 converts to XLVII", 47, "XLVII"},
		{"49 converts to XLIX", 49, "XLIX"},
		{"50 converts to L", 50, "L"},
		{"90 converts to XC", 90, "XC"},
		{"94 converts to XCIV", 94, "XCIV"},
		{"99 converts to XCIX", 99, "XCIX"},
		{"100 converts to C", 100, "C"},
		{"400 converts to CD", 400, "CD"},
		{"500 converts to D", 500, "D"},
		{"900 converts to CM", 900, "CM"},
		{"1000 converts to M", 1000, "M"},
		{"1984 converts to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 converts to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 converts to MMXIV", 2014, "MMXIV"},
		{"1006 converts to MVI", 1006, "MVI"},
		{"798 converts to DCCXCVIII", 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman((test.Num))
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}