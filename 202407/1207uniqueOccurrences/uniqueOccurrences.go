package main

func main() {

}

func uniqueOccurrences(arr []int) bool {

	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	for _, ch := range arr {
		cnt1[ch]++
	}
	for _, v := range cnt1 {
		cnt2[v]++
		if cnt2[v] > 1 {
			return false
		}
	}

	return true
}
