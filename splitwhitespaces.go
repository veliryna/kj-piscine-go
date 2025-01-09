package piscine

func SplitWhiteSpaces(s string) []string {
	var slice []string
	word := ""
	arrs := []rune(s)
	for i, value := range s {
		if string(value) != " " && string(value) != "\t" && string(value) != "\n" {
			word += string(value)
			if i == len(s)-1 {
				slice = append(slice, word)
			}
		} else {
			if i >= 1 && arrs[i-1] != ' ' && arrs[i-1] != '\t' && arrs[i-1] != '\n' {
				slice = append(slice, word)
			}
			word = ""
		}
	}
	return slice
}
