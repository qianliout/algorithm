package main

func main() {

}

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	ans := 0
	cnt := make([]int, n+10)
	sameCnt := 0
	mod := 0    // 最大的相同数
	modCnt := 0 // 最大的相同数的个数

	for i := range nums1 {
		x, y := nums1[i], nums2[i]
		if x == y {
			ans += i
			sameCnt++
			cnt[x]++
			if cnt[x] > modCnt {
				modCnt = cnt[x]
				mod = x
			}
		}
	}
	for i := range nums1 {
		x, y := nums1[i], nums2[i]
		if modCnt*2 <= sameCnt {
			break
		}
		if x != y && x != mod && y != mod {
			ans += i
			sameCnt++
		}
	}
	if modCnt*2 <= sameCnt {
		return int64(ans)
	}
	return -1
}
