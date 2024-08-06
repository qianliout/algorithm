package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cellsInRange("A1:F1"))
}

func cellsInRange(s string) []string {
	sp := strings.Split(s, ":")

	fir1, sec1 := split(sp[0])
	fir2, sec2 := split(sp[1])
	ans := make([]string, 0)
	for i := fir1; i <= fir2; i++ {
		for j := sec1; j <= sec2; j++ {
			ans = append(ans, string(i)+string(j))
		}
	}
	return ans
}

func split(s string) (byte, byte) {
	return s[0], s[1]
}
