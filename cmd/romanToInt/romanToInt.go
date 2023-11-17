package romanToInt

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	num := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		digit := romanMap[s[i]]
		if digit < prev {
			num -= digit
		} else {
			num += digit
		}
		prev = digit
	}
	return num
}
