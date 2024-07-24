package main

func main() {

}

func mergeTriplets(triplets [][]int, target []int) bool {
	ans := []int{0, 0, 0}
	for _, ch := range triplets {
		if canChoose(ch, target) {
			ans = choose(ch, ans)
		}
	}
	for i := 0; i < 3; i++ {
		if ans[i] != target[i] {
			return false
		}
	}
	return true
}

func choose(ch, ans []int) []int {
	for i := 0; i < 3; i++ {
		ans[i] = max(ans[i], ch[i])
	}
	return ans
}

func canChoose(ch, target []int) bool {
	for i := 0; i < 3; i++ {
		if ch[i] > target[i] {
			return false
		}
	}
	return true
}
