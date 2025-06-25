package main

func main() {

}

func subarrayBitwiseORs(arr []int) int {
	set := make(map[int]int)
	for i, ch := range arr {
		set[ch]++
		j := i - 1
		for j >= 0 && arr[j]|ch != arr[j] {
			arr[j] = arr[j] | ch
			set[arr[j]]++
			j--
		}
	}
	return len(set)
}
