package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
