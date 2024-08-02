package main

func main() {

}

func timeRequiredToBuy(tickets []int, k int) int {
	n := len(tickets)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		queue = append(queue, i)
	}
	ans := 0
	for tickets[k] > 0 {
		idx := queue[0]
		queue = queue[1:]
		tickets[idx]--
		ans++
		if tickets[idx] > 0 {
			queue = append(queue, idx)
		}
	}
	return ans
}
