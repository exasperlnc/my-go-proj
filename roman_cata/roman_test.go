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
		{"9 converts to XI", 9, "XI"},
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