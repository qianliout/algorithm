package main

func main() {

}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	le, ri := 0, len(plants)-1
	ans := 0
	ca, cb := capacityA, capacityB

	for le < ri {
		if plants[le] > ca && plants[le] <= capacityA {
			ca = capacityA
			ans++
		}
		if plants[ri] > cb && plants[ri] <= capacityB {
			cb = capacityB
			ans++
		}
		if ca >= plants[le] {
			ca -= plants[le]
			le++
		}
		if cb >= plants[ri] {
			cb -= plants[ri]
			ri--
		}
	}
	if le == ri {
		// 题目保证可以完成
		if plants[le] > max(ca, cb) {
			ans++
		}
	}
	return ans
}
