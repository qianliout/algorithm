package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
}

func complexNumberMultiply(num1 string, num2 string) string {
	// num1 = "1+1i", num2 = "1+1i"
	a, b := parse(num1)
	c, d := parse(num2)
	// (a+bi)(c+di)=(ac-bd)+(bc+ad)i。
	n := a*c - b*d
	m := b*c + a*d

	return fmt.Sprintf("%d+%di", n, m)
}

func parse(num string) (int, int) {
	split := strings.Split(num, "+")
	n, _ := strconv.Atoi(split[0])
	m1 := strings.ReplaceAll(split[1], "i", "")
	m, _ := strconv.Atoi(m1)
	return n, m
}
