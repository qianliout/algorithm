package main

func main() {

}

// nums 中的整数都有相同的数位长度。
func sumDigitDifferences(nums []int) int64 {
	ans := 0
	for {
		if nums[0] == 0 {
			break
		}
		cnt := make(map[int]int)
		for i, ch := range nums {
			a := ch % 10
			nums[i] = ch / 10
			ans += i - cnt[a]
			cnt[a]++
		}
	}
	return int64(ans)
}
