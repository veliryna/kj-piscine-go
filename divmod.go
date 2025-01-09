package piscine

func DivMod(a int, b int, div *int, mod *int) {
	value := a / b
	rem := a % b
	*div = value
	*mod = rem
}
