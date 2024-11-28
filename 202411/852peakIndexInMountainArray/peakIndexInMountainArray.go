package main

func main() {

}

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	le, ri := 0, n-1
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n-1 && arr[mid] > arr[mid+1] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}
