package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := false
	var s []int
	for i := 0; i < len(a)-1; i++ {
		s = append(s, f(a[i+1], a[i]))
	}
	countn, countp := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			countp++
		}
		if s[i] < 0 {
			countn--
		}
	}
	if countp == 0 || countn == 0 {
		res = true
	}
	return res
}
