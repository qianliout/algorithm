package main

func main() {

}

func canFormArray(arr []int, pieces [][]int) bool {
	idx := make(map[int]int)
	for i, ch := range arr {
		idx[ch] = i
	}
	n := len(arr)
	for _, p := range pieces {
		if len(p) == 0 {
			continue
		}
		first, ok := idx[p[0]]
		if !ok {
			return false
		}
		for j := first; j < first+len(p); j++ {
			if j >= len(arr) {
				return false
			}
			if arr[j] != p[j-first] {
				return false
			}
		}

		n -= len(p)
	}
	return n == 0
}
