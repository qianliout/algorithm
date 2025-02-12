package main

func main() {

}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	le, ri, n := 0, 0, len(customers)
	ans, win := 0, 0
	for i, ch := range grumpy {
		if ch == 0 {
			ans += customers[i]
		}
	}
	mx := 0
	for le <= ri && ri < n {
		if grumpy[ri] == 1 {
			win += customers[ri]
		}
		ri++
		mx = max(mx, win)
		if ri-le >= minutes {
			if grumpy[le] == 1 {
				win -= customers[le]
			}
			le++
		}
	}

	return ans + mx
}
