package main

import (
	"os"
)

const (
	INT_MAX = 9223372036854775807
	INT_MIN = -9223372036854775808
)

func IsNum(s string) bool {
	res := true
	for index, value := range s {
		if index == 0 && value == '-' {
			continue
		}
		if value >= '0' && value <= '9' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}

func ToInt(s string) int {
	res := 0
	counter := 1
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		if i == 0 && r[i] == '-' {
			res = -1 * res
		} else {
			res += int(r[i]-'0') * counter
			counter *= 10
		}
	}
	return res
}

func ToString(s int) string {
	res := ""
	neg := false
	if s == 0 {
		res = "0"
	}
	if s < 0 {
		s = -1 * s
		neg = true
	}
	for s > 0 {
		d := s % 10
		res = string(rune(d+'0')) + res
		s = s / 10
	}
	if neg {
		res = "-" + res
	}
	return res
}

func main() {
	arg := os.Args
	if len(arg) == 4 {
		value1 := arg[1]
		op := arg[2]
		value2 := arg[3]
		if IsNum(value1) && IsNum(value2) {
			int1 := ToInt(value1)
			int2 := ToInt(value2)
			var result int
			switch op {
			case "+":
				if !((int2 > 0 && int1 > INT_MAX-int2) || (int2 < 0 && int1 < INT_MIN-int2)) {
					result = int1 + int2
					os.Stdout.Write([]byte(ToString(result)))
					os.Stdout.Write([]byte("\n"))
				}
			case "-":
				if !((int2 < 0 && int1 > INT_MAX+int2) || (int2 > 0 && int1 < INT_MIN+int2)) {
					result = int1 - int2
					os.Stdout.Write([]byte(ToString(result)))
					os.Stdout.Write([]byte("\n"))
				}
			case "*":
				if !((int2 != 0 && int1 > INT_MAX/int2) || (int2 != 0 && int1 < INT_MIN/int2)) {
					result = int1 * int2
					os.Stdout.Write([]byte(ToString(result)))
					os.Stdout.Write([]byte("\n"))
				}
			case "/":
				if int2 != 0 {
					result = int1 / int2
					os.Stdout.Write([]byte(ToString(result)))
					os.Stdout.Write([]byte("\n"))
				} else {
					os.Stdout.Write([]byte("No division by 0"))
					os.Stdout.Write([]byte("\n"))
				}
			case "%":
				if int2 != 0 {
					result = int1 % int2
					os.Stdout.Write([]byte(ToString(result)))
					os.Stdout.Write([]byte("\n"))
				} else {
					os.Stdout.Write([]byte("No modulo by 0"))
					os.Stdout.Write([]byte("\n"))
				}
			}
		}
	}
}
