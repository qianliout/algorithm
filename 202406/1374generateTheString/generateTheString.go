package main

func main() {

}

func generateTheString(n int) string {
	nums := h(n)
	ans := make([]byte, 0)
	start := 'a'
	for i := range nums {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, byte(start))
		}
		start++
	}
	return string(ans)
}

func h(n int) []int {
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}
	if n%2 == 0 {
		return []int{1, n - 1}
	} else {
		return []int{1, 1, n - 2}
	}
}
