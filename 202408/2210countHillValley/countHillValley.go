package main

func main() {

}

func countHillValley(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if check1(nums, i) || check2(nums, i) {
			ans++
		}
	}
	return ans
}

// 波谷
func check1(nums []int, i int) bool {
	n := len(nums)
	cnt := 0
	for j := i - 1; j >= 0; j-- {
		if nums[j] == nums[i] {
			continue
		} else {
			if nums[j] > nums[i] {
				cnt++
			}
			break
		}
	}
	for j := i + 1; j < n; j++ {
		if nums[j] == nums[i] {
			continue
		} else {
			if nums[j] > nums[i] {
				cnt++
			}
			break
		}
	}

	return cnt == 2
}
func check2(nums []int, i int) bool {
	n := len(nums)
	cnt := 0
	for j := i - 1; j >= 0; j-- {
		if nums[j] == nums[i] {
			continue
		} else {
			if nums[j] < nums[i] {
				cnt++
			}
			break
		}
	}
	for j := i + 1; j < n; j++ {
		if nums[j] == nums[i] {
			continue
		} else {
			if nums[j] < nums[i] {
				cnt++
			}
			break
		}
	}

	return cnt == 2
}
