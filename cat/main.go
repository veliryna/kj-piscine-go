package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	err := "ERROR: "
	runes := []rune(err)
	arg := os.Args[1:]
	le := 0
	for i := range arg {
		le = i + 1
	}
	if le == 0 {
		input, er := ioutil.ReadAll(os.Stdin)
		for _, j := range string(input) {
			z01.PrintRune(j)
		}
		if er != nil {
			for _, k := range er.Error() {
				z01.PrintRune(k)
			}
			z01.PrintRune(10)
		}
		return
	}
	one := true
	for _, ar := range arg {
		file, er := os.Open(ar)
		if er != nil {
			for _, r := range runes {
				z01.PrintRune(r)
			}
			for _, e := range er.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune(10)
			os.Exit(1)
		}
		k, er := ioutil.ReadAll(file)
		if !one {
		}
		one = false
		for _, t := range string(k) {
			z01.PrintRune(t)
		}
		if er != nil {
			for _, e := range er.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune(10)
		}
		file.Close()
	}
}
