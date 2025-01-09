package piscine

func Any(f func(string) bool, a []string) bool {
	res := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			res = true
		}
	}
	return res
}
