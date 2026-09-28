package piscine

func ConcatParams(args []string) string {
	res := ""
	for i, char := range args {
		res += char
		if i != len(args)-1 {
			res += "\n"
		}
	}
	return res
}
