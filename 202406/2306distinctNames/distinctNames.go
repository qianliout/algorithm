package main

func main() {

}

func distinctNames(ideas []string) int64 {
	group := make(map[string]int)
	for _, s := range ideas {
		group[s[1:]] |= 1 << (int(s[0]) - int('a'))
	}

	/*
		义 cnt[i][j] 表示组中首字母不包含 i 但包含 j 的组的个数。枚举每个组，统计 cnt，
		同时枚举该组的首字母 i 和 不在该组的首字母 j，答案即为 cnt[i][j] 的累加值。
		简单来说就是「有 i 无 j」可以和「无 i 有 j」的字符串互换。
		由于我们是一次遍历所有组，没有考虑两个字符串的顺序，最后需要把答案乘 2，表示 A+B 和 B+A 两种字符串的组合。
		所以只需要计算存在 i 但是不存在 j 的数据就行
	*/
	cnt := [26][26]int{}
	var ans int64
	for _, mask := range group {
		for i := 0; i < 26; i++ {
			// 不存在i
			if mask>>i&1 == 0 {
				for j := 0; j < 26; j++ {
					if mask>>j&1 == 1 {
						// 不存在 i 但是存在 j
						cnt[i][j]++
					}
				}
			} else {
				// 存在 i 累加不存在 j 的数据就是答案
				for j := 0; j < 26; j++ {
					if mask>>j&1 == 0 {
						ans += int64(cnt[i][j])
					}
				}
			}
		}
	}
	return 2 * ans
}
