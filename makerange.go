package piscine

func MakeRange(min, max int) []int {
	if min < max {
		s := make([]int, max-min)
		for i := 0; i < len(s); i++ {
			s[i] = min
			min++
		}
		return s
	} else {
		var s []int
		return s
	}
}
