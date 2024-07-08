package main

func main() {

}

// 未解决
type MajorityChecker struct {
	Data map[int]int
}

func Constructor(arr []int) MajorityChecker {
	cnt := make(map[int]int)
	for _, ch := range arr {
		cnt[ch]++
	}
	return MajorityChecker{Data: cnt}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for i := left; i <= right; i++ {
		if this.Data[i] >= threshold {
			return i
		}
	}
	return -1
}
