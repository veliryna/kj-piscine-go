package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return false
	}
	a := []rune(s)
	for _, letter := range a {
		if letter >= 32 && letter <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
