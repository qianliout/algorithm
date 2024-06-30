package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
}

func discountPrices(sentence string, discount int) string {
	split := strings.Split(sentence, " ")
	for i, ch := range split {
		n, b := check(ch)
		if b {
			split[i] = fmt.Sprintf("$%s", parse(n, discount))
		}
	}
	return strings.Join(split, " ")

}

func check(s string) (int, bool) {
	if s[0] != '$' {
		return 0, false
	}
	if len(s) > 11 {
		return 0, false
	}

	atoi, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, false
	}
	return atoi, true
}

func parse(a int, dis int) string {
	return fmt.Sprintf("%.2f", (float64(a*(100-dis)))/float64(100))
}
