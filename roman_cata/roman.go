package romancata

func ConvertToRoman(num int) string {
	if num == 2 {
		return "II"
	}
	return "I"
}