package main

func main() {

}

func numberOfWays(s string) int64 {
	ans := 0
	n0 := 0 // 子序列 0
	n1 := 0
	n10 := 0 // 10
	n01 := 0 // 01
	for _, ch := range s {
		if byte(ch) == '1' {
			n1++
			n01 = n01 + n0
			ans += n10
		} else {
			n0++
			n10 = n10 + n1
			ans += n01
		}
	}
	return int64(ans)
}
