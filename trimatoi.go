package piscine

func TrimAtoi(s string) int {
	result := ""
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] == '-' && result == "" {
			result += (string(str[i]))
		}
		if str[i] >= '0' && str[i] <= '9' {
			result += string(str[i])
		}
	}
	num := 0
	negative := false
	for _, digit := range result {
		if digit == '-' {
			negative = true
			continue
		}
		num = num*10 + int(digit-'0')
	}
	if negative {
		num = -num
	}
	return num
}
