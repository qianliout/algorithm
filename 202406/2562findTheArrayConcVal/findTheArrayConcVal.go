package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func findTheArrayConcVal(nums []int) int64 {
	ans := 0
	le, ri := 0, len(nums)-1
	for le <= ri {
		if le == ri {
			ans += nums[le]
			break
		} else {
			to, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[le], nums[ri]))
			ans += to
			le++
			ri--
		}
	}
	return int64(ans)
}
