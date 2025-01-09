package piscine

func IterativeFactorial(nb int) int {
	var res int
	if nb < 0 || nb > 20 {
		res = 0
	} else if nb == 0 {
		res = 1
	} else {
		res = 1
		for i := 1; i <= nb; i++ {
			res *= i
		}
	}
	return res
}
