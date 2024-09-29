package main

import "fmt"

func main() {
	fmt.Println(validStrings(3))
}

func validStrings(n int) []string {
	mask := 1<<n - 1
	ans := make([]string, 0)
	for i := 0; i < 1<<n; i++ {
		x := i ^ mask
		if x&(x>>1) == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, i))
		}
	}
	return ans
}
