package main

import (
	"math/bits"
)

func main() {

}

func canMakePaliQueries1(s string, queries [][]int) []bool {
	sum := make([][26]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		sum[i+1] = sum[i]
		sum[i+1][s[i]-'a']++
	}
	ans := make([]bool, len(queries))
	for i, qu := range queries {
		le, ri, k, m := qu[0], qu[1], qu[2], 0
		for j := 0; j < 26; j++ {
			// 这里容易出错的点是 sum[ri+1]
			if (sum[ri+1][j]-sum[le][j])%2 != 0 {
				m++
			}
		}
		ans[i] = m/2 <= k
	}
	return ans
}

// 只关心奇偶性
func canMakePaliQueries2(s string, queries [][]int) []bool {
	sum := make([][26]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		sum[i+1] = sum[i]
		sum[i+1][s[i]-'a'] ^= 1
	}
	ans := make([]bool, len(queries))
	for i, qu := range queries {
		le, ri, k, m := qu[0], qu[1], qu[2], 0
		for j := 0; j < 26; j++ {
			// 这里容易出错的点是 sum[ri+1]
			// 这样写也是对的,因为 sum[ri+1][j]里的值要么是1，要么是0
			//  if sum[ri+1][j]!=sum[le][j] {
			if sum[ri+1][j]^sum[le][j] == 1 {
				m++
			}
		}
		ans[i] = m/2 <= k
	}
	return ans
}

// 状态压缩
// 因为只有26个字母，所以让26个位分别表示每个字母
func canMakePaliQueries(s string, queries [][]int) []bool {
	sum := make([]uint32, len(s)+1)
	for i := 0; i < len(s); i++ {
		sum[i+1] = sum[i] ^ (1 << (s[i] - 'a'))
	}
	ans := make([]bool, len(queries))
	for i, qu := range queries {
		le, ri, k, m := qu[0], qu[1], qu[2], 0
		m += bits.OnesCount32(sum[ri+1] ^ sum[le])
		ans[i] = m/2 <= k
	}
	return ans
}
