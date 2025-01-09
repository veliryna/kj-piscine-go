package piscine

func IsUpper(str string) bool {
	check := true
	for _, value := range str {
		if value >= 'A' && value <= 'Z' {
			check = check
		} else {
			check = false
		}
	}
	return check
}
