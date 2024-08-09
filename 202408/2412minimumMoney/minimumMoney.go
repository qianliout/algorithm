package main

func main() {

}

func minimumMoney(transactions [][]int) int64 {
	allLost := 0
	mx := 0
	for _, ch := range transactions {
		cost, cashback := ch[0], ch[1]
		allLost += max(cost-cashback, 0)
		mx = max(mx, min(cost, cashback))
	}
	return int64(allLost + mx)
}
