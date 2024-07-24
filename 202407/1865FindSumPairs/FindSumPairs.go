package main

func main() {

}

type FindSumPairs struct {
	Nums1 []int
	Nums2 []int
	Cnt   map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	cnt := make(map[int]int)
	for _, ch := range nums2 {
		cnt[ch]++
	}
	f := FindSumPairs{
		Nums1: nums1,
		Nums2: nums2,
		Cnt:   cnt,
	}
	return f

}

func (this *FindSumPairs) Add(index int, val int) {
	pre := this.Nums2[index]
	this.Nums2[index] += val
	this.Cnt[pre]--
	this.Cnt[pre+val]++
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for _, ch := range this.Nums1 {
		ans += this.Cnt[tot-ch]
	}
	return ans
}
