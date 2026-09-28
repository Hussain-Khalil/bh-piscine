package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	output := ""
	for i, chr := range str {
		if (i+1)%3 == 0 {
			output += string(chr)
		}
	}
	if len(output) == 0 {
		return "\n"
	}
	return output + "\n"
}
