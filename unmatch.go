package piscine

func Unmatch(a []int) int {
	for _, v := range a {
		n := 0
		for _, c := range a {
			if c == v {
				n++
			}
		}
		if n == 1 || n%2 == 1 {
			return v
		}
	}
	return -1
}
