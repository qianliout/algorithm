package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("12936"))
	fmt.Println(countSubstrings("5701283"))
	fmt.Println(countSubstrings("1010101010"))
}

func countSubstrings(s string) int64 {
	ans := 0
	f := make([][]int, 10)
	for i := range f {
		f[i] = make([]int, 10)
	}
	for _, ch := range s {
		d := int(ch) - int('0')
		//  不能加这一个剪枝，为啥呢？
		// if d == 0 {
		// 	continue
		// }
		for m := 1; m <= 9; m++ {
			nf := make([]int, 10)
			nf[d%m] = 1
			for rem := 0; rem < m; rem++ {
				nf[(rem*10+d)%m] += f[m][rem]
			}
			f[m] = nf
		}
		ans += f[d][0]
	}
	return int64(ans)
}
