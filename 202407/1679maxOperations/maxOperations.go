package main

func main() {

}

func maxOperations(nums []int, k int) int {
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch]++
	}
	ans := 0
	for key, v := range cnt {

		if v == 0 {
			continue
		}

		nex := k - key
		// 这个特判容易出错
		if key == nex {
			ans += v / 2
			cnt[key] -= v / 2
			continue
		}
		b := min(v, cnt[k-key])
		ans += b
		cnt[key] -= b
		cnt[k-key] -= b
	}
	return ans
}
