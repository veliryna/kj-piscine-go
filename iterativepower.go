package piscine

func IterativePower(nb int, power int) int {
	var res int
	if power < 0 {
		res = 0
	} else if power == 0 {
		res = 1
	} else {
		res = nb
		for i := 2; i <= power; i++ {
			res *= nb
		}
	}
	return res
}
