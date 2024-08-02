package main

import (
	"fmt"
)

func main() {
	fmt.Println(wateringPlants([]int{2, 2, 3, 3}, 5))
}

func wateringPlants(plants []int, capacity int) int {
	ans := 0
	i := 0
	n := len(plants)
	k := capacity
	for i < n {
		ch := plants[i]
		if k >= ch {
			k -= ch
			i++
			ans++
		} else {
			ans += 2 * i
			k = capacity
		}
	}
	return ans
}

/*
按从左到右的顺序给植物浇水。
在给当前植物浇完水之后，如果你没有足够的水 完全 浇灌下一株植物，那么你就需要返回河边重新装满水罐。
你 不能 提前重新灌满水罐。
*/
