package main

func main() {

}

func longestMountain(arr []int) int {
	n := len(arr)
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			suf[i] = suf[i+1] + 1
		}
	}

	pre := make([]int, n)
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			pre[i] = pre[i-1] + 1
		}
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if pre[i] > 0 && suf[i] > 0 {
			ans = max(ans, pre[i]+suf[i]+1)
		}
	}
	return ans
}

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	if n < 3 {
		return -1
	}
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			suf[i] = suf[i+1] + 1
		}
	}

	pre := make([]int, n)
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			pre[i] = pre[i-1] + 1
		}
	}
	ans := 0
	idx := -1
	for i := 1; i < n-1; i++ {
		if pre[i] > 0 && suf[i] > 0 && pre[i]+suf[i]+1 > ans {
			ans = max(ans, pre[i]+suf[i]+1)
			idx = i
		}
	}
	return idx
}
