package main

import "fmt"

func main() {
	fmt.Println(balancedString("QWER"))
	fmt.Println(balancedString("WQWRQQQW"))
}

func balancedString(s string) int {
	cntQ, cntW, cntE, cntR, n := 0, 0, 0, 0, len(s)
	for _, ch := range s {
		switch ch {
		case 'Q':
			cntQ++
		case 'W':
			cntW++
		case 'E':
			cntE++
		case 'R':
			cntR++
		}
	}
	if cntQ <= n/4 && cntW <= n/4 && cntE <= n/4 && cntR <= n/4 {
		return 0
	}
	le, ri := 0, 0
	ans := n + 1
	for le <= ri && ri < n {
		switch s[ri] {
		case 'Q':
			cntQ--
		case 'W':
			cntW--
		case 'E':
			cntE--
		case 'R':
			cntR--
		}
		for le <= ri && check(cntQ, cntW, cntE, cntR, n) {
			ans = min(ans, ri-le+1)
			switch s[le] {
			case 'Q':
				cntQ++
			case 'W':
				cntW++
			case 'E':
				cntE++
			case 'R':
				cntR++
			}
			le++
		}
		ri++

	}
	return ans
}

func check(cntQ, cntW, cntE, cntR, n int) bool {
	if cntQ <= n/4 && cntW <= n/4 && cntE <= n/4 && cntR <= n/4 {
		return true
	}
	return false
}
