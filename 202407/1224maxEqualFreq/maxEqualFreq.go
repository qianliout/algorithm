package main

func main() {

}

func maxEqualFreq(nums []int) int {
	cnt := make(map[int]int)
	sum := make(map[int]int)
	ans := 0
	mx := 0
	for i, ch := range nums {
		// cnt用于记录数组中每个元素的出现频率；
		// sum用于记录相同频率的元素个数；
		// ans用于记录满足条件的最长子数组长度；
		// mx用于记录当前最大频率。
		cnt[ch]++
		cur := cnt[ch]
		sum[cur]++
		sum[cur-1]--
		mx = max(mx, cur)
		// 如果最大频率mx为1，说明数组中只有一个元素，此时该子数组满足要求，更新ans为当前位置加1
		if mx == 1 {
			ans = max(ans, i+1)
		}
		// 如果最大频率mx乘以sum[mx]加1等于当前位置加1，说明可以通过移除一个频率最高的元素，使得剩余元素频率相等，更新ans为当前位置加1。
		if mx*sum[mx]+1 == (i + 1) {
			ans = max(ans, i+1)
		}

		// (max−1)×(sum[max−1]+1)+1=len：说明出现次数为 max 的数值只有一个，其余出现次数均为 max - 1，对其删除一次后即满足要求
		// （删除该出现次数为 max 的数值后，会导致出现次数为 max - 1 的数值多一个，此时有「出现次数为 max - 1 的数值 + 被删除的一个数 = 总数量 len」）

		if (mx-1)*(sum[mx-1]+1)+1 == i+1 {
			ans = max(ans, i+1)
		}
	}
	return ans
}
