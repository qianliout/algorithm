package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(smallestNumber(-7605))
	fmt.Println(uint64(1<<64 - 1))
	fmt.Println(uint64(math.MaxUint64))
	fmt.Println(bits.OnesCount64(math.MaxUint64))
	fmt.Println(bits.OnesCount64(uint64(1 << 63)))
}

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}
	pre := num

	nums, zero := make([]int, 0), make([]int, 0)

	for num != 0 {
		a := num % 10
		num = num / 10
		if a == 0 {
			zero = append(zero, 0)
		} else {
			nums = append(nums, abs(int(a)))
		}
	}
	if pre < 0 {
		return int64(-1 * genMx(nums, zero))
	}
	return int64(genMin(nums, zero))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func genMx(nums []int, zero []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	nums = append(nums, zero...)
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans*10 + nums[i]
	}
	return ans
}

func genMin(nums []int, zero []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	a := []int{nums[0]}
	a = append(a, zero...)
	a = append(a, nums[1:]...)
	ans := 0
	for i := 0; i < len(a); i++ {
		ans = ans*10 + a[i]
	}
	return ans
}

//-1015 <= num <= 1015
