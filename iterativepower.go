package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || nb > 20 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	result := nb
	for i := 1; i < power; i++ {
		result *= nb
	}
	return result
}
