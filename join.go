package piscine

func Join(strs []string, sep string) string {
	output := ""
	for i, chr := range strs {
		if i == len(strs)-1 {
			output += chr
		} else {
			output += chr
			output += sep

		}
	}
	return output
}
