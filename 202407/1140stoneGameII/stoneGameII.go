package main

func main() {

}

func stoneGameII(piles []int) int {
	var dfs func(nums []int, m int) (int, int)

	dfs = func(nums []int, m int) (int, int) {
		if len(nums) == 0 {
			return 0, 0
		}
		res := 0
		sum := nums[0]
		for i := 1; i <= 2*m; i++ {

			res = max(res)

		}

	}

}
