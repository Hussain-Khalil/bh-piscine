package piscine

func Rot14(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			rot := (int(char) - 'a' + 14) % 26
			result += string(rot + 'a')
		} else if char >= 'A' && char <= 'Z' {
			rot := (int(char) - 'A' + 14) % 26
			result += string(rot + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
