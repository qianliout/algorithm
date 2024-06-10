package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(divisorSubstrings(430043, 2))
	fmt.Println(divisorSubstrings2(430043, 2))
}

func divisorSubstrings(pre int, k int) int {
	wind, ans := make([]byte, 0), 0
	num := []byte(strconv.Itoa(pre))
	ri := 0

	for ri < len(num) {
		wind = append(wind, num[ri])
		ri++
		for len(wind) > k {
			wind = wind[1:]
		}

		if len(wind) == k {
			o, _ := strconv.Atoi(string(wind))
			if o != 0 && pre%o == 0 {
				ans++
			}
		}
	}
	return ans
}

func divisorSubstrings2(num, k int) int {
	ans := 0
	s := strconv.Itoa(num)

	for i := 0; i+k <= len(s); i++ {
		v, _ := strconv.Atoi(s[i : i+k])
		if v != 0 && num%v == 0 {
			ans++
		}
	}
	return ans
}

func divisorSubstrings3(num, k int) (ans int) {
	m := int(math.Pow10(k))
	for x := num; x >= m/10; x /= 10 {
		if x%m > 0 && num%(x%m) == 0 {
			ans++
		}
	}
	return
}

func revers(num int) int {
	after := 0
	for num > 0 {
		after = after*10 + num%10
		num = num / 10
	}
	return after
}
