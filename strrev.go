package piscine

func StrRev(s string) string {
	arr := []rune(s)
	nm := []rune{}
	for i := len(arr) - 1; i >= 0; i-- {
		nm = append(nm, arr[i])
	}
	return string(nm)
}
