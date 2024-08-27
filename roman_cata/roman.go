package romancata

import "strings"

func ConvertToRoman(num int) string {
	var result strings.Builder
	
	for i := num; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}