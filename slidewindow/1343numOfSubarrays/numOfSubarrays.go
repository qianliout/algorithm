package main

func main() {

}

func numOfSubarrays(arr []int, k int, threshold int) int {
	le, ri, n := 0, 0, len(arr)
	ans := 0
	win := 0

	for le <= ri && ri < n {
		win += arr[ri]
		ri++
		if ri-le >= k {
			if win >= threshold*k {
				ans++
			}
		}
		if ri-le >= k {
			win -= arr[le]
			le++
		}
	}
	return ans
}
