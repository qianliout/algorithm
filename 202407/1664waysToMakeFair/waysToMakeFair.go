package main

func main() {

}

// 你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。
func waysToMakeFair(nums []int) int {
	s1 := 0 // 偶数下标前缀和
	s2 := 0 // 奇数下标前缀和
	t1 := 0 // 遍历 i ，对于 i 前 偶数下标前缀和
	t2 := 0 // 遍历 i ，对于 i 前 奇数下标前缀和

	ans := 0
	for i, v := range nums {
		if i&1 == 0 {
			s1 += v
		} else {
			s2 += v
		}
	}
	for i, v := range nums {
		// 假设删除 i
		if i&1 == 0 {
			// 如果删除 i，此时 i 前面的小标的奇偶性不变，后面的偶数下标是奇数，奇数下标是偶数
			if t2+(s1-t1-v) == t1+(s2-t2) {
				ans++
			}
		}
		if i&1 == 1 {
			if t2+s1-t1 == t1+s2-t2-v {
				ans++
			}
		}
		if i&1 == 1 {
			t2 += v
		} else {
			t1 += v
		}
	}

	return ans
}
