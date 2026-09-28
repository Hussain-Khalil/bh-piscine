package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := false
	if len(a) == 1 {
		return true
	} else if len(a) < 1 {
		return false
	}
	if f(a[0], a[1]) <= 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) <= 0 {
				arr = true
			} else {
				return false
			}
		}
	} else if f(a[0], a[1]) >= 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) >= 0 {
				arr = true
			} else {
				return false
			}
		}
	}
	return arr
}
