package main

func main() {

}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// 其实就是求 arr1中有多少个元素满足要求，数据量不大，直接暴力解决
	cnt := 0
	for _, ch1 := range arr1 {
		flag := true
		for _, ch2 := range arr2 {
			if abs(ch1-ch2) <= d {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	return cnt
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
