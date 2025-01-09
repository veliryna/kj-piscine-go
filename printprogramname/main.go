package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	arr := []rune(arguments[0])
	for _, r := range arr {
		if r != '.' && r != '/' {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
