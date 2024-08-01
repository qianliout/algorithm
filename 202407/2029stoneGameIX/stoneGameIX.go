package main

func main() {

}

func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)
	for _, ch := range stones {
		cnt[ch%3]++
	}
	if cnt[0]%2 == 0 {
		return !(cnt[1] == 0 || cnt[2] == 0)
	}
	return abs(cnt[1]-cnt[2]) > 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
