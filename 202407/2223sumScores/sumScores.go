package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumScores("azbazbzaz"))
}

func sumScores(s string) int64 {

	n := len(s)
	hash, p := preHash(s)

	var ans int64 = 0
	for i := 0; i < n; i++ {
		ans += cal2(i, n, hash, p)
	}

	return ans
}

// 计算最长公共前缀
// 直接计算会超时，只能是用二分
func cal(j, n int, hash, p []int64) int64 {
	var ans int
	l := n - j
	mod := int64(math.Pow10(9)) + 7
	for i := 1; i <= l; i++ {
		le := hash[i]
		ri := (hash[j+i] - hash[j]*p[i]%mod + mod) % mod
		if le == ri {
			ans = max(ans, i)
		}
	}
	return int64(ans)
}

func cal2(j, n int, hash, p []int64) int64 {
	l := n - j
	mod := int64(math.Pow10(9)) + 7
	le, ri := 0, l+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid > n || mid < 0 || j+mid > n {
			break
		}
		lef := hash[mid]
		rig := (hash[j+mid] - hash[j]*p[mid]%mod + mod) % mod
		if lef == rig {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	return int64(le)
}

func preHash(text string) ([]int64, []int64) {
	n := len(text)
	hash := make([]int64, n+1)
	p := make([]int64, n+1)
	p[0] = 1
	var base int64 = 27
	mod := int64(math.Pow10(9)) + 7
	for i, ch := range text {
		hash[i+1] = (hash[i]*base + encode(byte(ch))) % mod
		p[i+1] = p[i] * base % mod
	}
	return hash, p
}

func encode(ch byte) int64 {
	return int64(ch) - int64('a') + 1
}
