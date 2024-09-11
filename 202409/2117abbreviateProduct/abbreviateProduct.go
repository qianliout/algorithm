package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

}

func abbreviateProduct(left, right int) string {
	s := new(big.Int).MulRange(int64(left), int64(right)).String()
	tz := len(s)
	s = strings.TrimRight(s, "0")
	tz -= len(s)
	if len(s) > 10 {
		return fmt.Sprintf("%s...%se%d", s[:5], s[len(s)-5:], tz)
	}
	return fmt.Sprintf("%se%d", s, tz)
}
