package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	g := 0
	allGas := 0
	allCost := 0
	n := len(gas)
	for i := 0; i < n; i++ {
		allGas += gas[i]
		allCost += cost[i]
	}
	if allGas < allCost {
		return -1
	}
	for i := 0; i < n; i++ {
		if g < 0 {
			start = i
			g = 0
		}
		g = g + gas[i] - cost[i]
	}
	return start
}
