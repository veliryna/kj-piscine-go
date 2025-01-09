package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i >= 1; i-- {
		arr := []rune(arguments[i])
		for _, r := range arr {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
