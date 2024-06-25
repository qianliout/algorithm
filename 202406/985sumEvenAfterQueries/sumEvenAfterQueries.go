package main

func main() {

}

func sumEvenAfterQueries2(nums []int, queries [][]int) []int {
	sum := 0
	for _, ch := range nums {
		if ch%2 == 0 {
			sum += ch
		}
	}
	ans := make([]int, len(nums))
	for i := range queries {
		// 这样为啥会错呢，因为修改是永久修改，后面可能会有重复的查询，如果有重复的查询就会出错
		add, inx := queries[i][0], queries[i][1]
		added := nums[inx] + add
		if nums[inx]%2 == 0 {
			sum -= nums[inx]
		}
		if added%2 == 0 {
			sum += added
		}
		ans[i] = sum
	}
	return ans
}

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sum := 0
	for _, ch := range nums {
		if ch%2 == 0 {
			sum += ch
		}
	}
	ans := make([]int, len(nums))
	for i := range queries {
		add, inx := queries[i][0], queries[i][1]
		added := nums[inx] + add
		if nums[inx]%2 == 0 {
			sum -= nums[inx]
		}
		// 为啥要修改 nums[inx]呢
		// 因为修改是永久修改，后面可能会有重复的查询，如果有重复的查询就会出错
		nums[inx] = added
		if nums[inx]%2 == 0 {
			sum += nums[inx]
		}
		ans[i] = sum
	}
	return ans
}
