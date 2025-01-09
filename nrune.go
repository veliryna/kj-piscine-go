package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n <= 0 {
		return '\x00'
	}
	r := []rune(s)
	return r[n-1]
}
