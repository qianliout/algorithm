package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func convertDateToBinary1(date string) string {
	split := strings.Split(date, "-")
	ans := make([]string, 0)
	for _, ch := range split {
		n, _ := strconv.Atoi(ch)
		b := fmt.Sprintf("%b", n)
		ans = append(ans, b)
	}
	return strings.Join(ans, "-")
}

func convertDateToBinary(date string) string {
	split := strings.Split(date, "-")
	ans := make([]string, 0)
	for _, ch := range split {
		n, _ := strconv.Atoi(ch)
		b := strconv.FormatInt(int64(n), 2)
		ans = append(ans, b)
	}
	return strings.Join(ans, "-")
}
