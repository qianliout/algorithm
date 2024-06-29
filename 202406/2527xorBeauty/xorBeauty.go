package main

func main() {

}

/*
巧妙的理解方式
((nums[i] | nums[j]) & nums[k])
0 <= i, j, k < n
根据取值方式，如果第一次取 (nums[i] | nums[j]) 那么下一次就可以取(nums[j] | nums[i]) 如果i!=j 那么结果是0，所以只有当 j==i 时才会对
结果有贡献，也就是求 （nums[i]&nums[k]）
同样（nums[i]&nums[k]）^(nums[k]&nums[i]）) 也只有当k==i 时才会对答案有贡献
*/
func xorBeauty(nums []int) int {
	res := 0
	for _, ch := range nums {
		res = res ^ ch
	}
	return res
}
