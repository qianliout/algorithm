package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minMaxDifference(11891))
}

func minMaxDifference(num int) int {
	mx := []byte(strconv.Itoa(num))
	mn := []byte(strconv.Itoa(num))

	for i := 0; i < len(mx); i++ {
		if mx[i] != '9' {
			pre := mx[i]
			for j := i; j < len(mx); j++ {
				if mx[j] == pre {
					mx[j] = '9'
				}
			}
			break
		}
	}

	for i := 0; i < len(mn); i++ {
		if mn[i] != '0' {
			pre := mn[i]
			for j := i; j < len(mn); j++ {
				if mn[j] == pre {
					mn[j] = '0'
				}
			}
			break
		}
	}
	mx1, _ := strconv.Atoi(string(mx))
	mm1, _ := strconv.Atoi(string(mn))
	return mx1 - mm1
}
