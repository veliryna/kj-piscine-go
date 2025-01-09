package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func checkPos(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]
		if other_row_pos == row_position ||
			other_row_pos == row_position-(queen_number-i) ||
			other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func getSolution(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			var r rune
			r = rune(position[i] + 1 + 48)
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if checkPos(k, i) {
				position[k] = i
				getSolution(k + 1)
			}
		}
	}
}

func EightQueens() {
	getSolution(0)
}
