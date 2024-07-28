package main

func main() {

}

func maxCount(banned []int, n int, maxSum int) int {
	m := 10001
	exist := make([]int, m)
	for _, ch := range banned {
		exist[ch]++
	}
	all := 0
	ans := 0
	for i := 1; i <= n; i++ {
		if exist[i] > 0 {
			continue
		}
		if all+i > maxSum {
			break
		}
		all += i
		ans++
	}
	return ans
}
