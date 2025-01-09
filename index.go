package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	if len(toFind) == 0 {
		return 0
	}
	if len(toFind) > len(s) {
		return -1
	}
	if len(toFind) == len(s) && s == toFind {
		return 0
	}
	if len(toFind) == 1 {
		for index, value := range a {
			if value == b[0] {
				return index
			}
		}
	}
	if len(toFind) > 1 && len(toFind) < len(s) {
		check := true
		for index, value := range a {
			if value == b[0] {
				myIndex := index
				for _, value2 := range b {
					if value2 != a[index] {
						check = false
						break
					}
					index++
				}
				if check {
					return myIndex
				}
				check = true
			}
		}
	}
	return -1
}
