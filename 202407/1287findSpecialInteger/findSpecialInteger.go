package main

func main() {

}

func findSpecialInteger(arr []int) int {
	start := 0
	n := len(arr)
	for start < n {
		cnt := 0
		j := start
		for ; j < n; j++ {
			if arr[j] != arr[start] {
				break
			}
			cnt++
		}
		if cnt > n/4 {
			return arr[start]
		}
		start = j
	}

	return -1
}
