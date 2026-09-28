package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	number := 0
	for _, c := range nbr {
		digit := 0
		for i, b := range baseFrom {
			if c == b {
				digit = i
				break
			}
		}
		number = number*len(baseFrom) + digit
	}
	if number == 0 {
		return string(baseTo[0])
	}
	result := ""
	for number > 0 {
		remainder := number % len(baseTo)
		result = string(baseTo[remainder]) + result
		number /= len(baseTo)
	}
	return result
}
