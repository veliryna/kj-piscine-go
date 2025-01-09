package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printCoords(k int) {
	var slice []int
	for k > 0 {
		d := k % 10
		slice = append([]int{d}, slice...)
		k = k / 10
	}
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(slice[i] + 48))
	}
}

func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printCoords(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printCoords(points.y)
	z01.PrintRune('\n')
}
