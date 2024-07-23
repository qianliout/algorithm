package main

func main() {

}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	cnt := make(map[int]map[int]int)
	for _, ch := range logs {
		id, ti := ch[0], ch[1]
		if cnt[id] == nil {
			cnt[id] = make(map[int]int)
		}
		cnt[id][ti]++
	}

	// 下标从1开始
	ans := make([]int, k+1)
	for _, v := range cnt {
		ans[len(v)]++
	}

	return ans[1:]
}
