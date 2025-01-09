package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BubbleSort(array []string) []string {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := []string(os.Args[1:])
	BubbleSort(array)
	for i := 1; i < len(array); i++ {
		arr := []rune(array[i])
		for _, r := range arr {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
