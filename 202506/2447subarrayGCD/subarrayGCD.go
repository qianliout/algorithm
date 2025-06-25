package main

func main() {

}

func subarrayGCD(nums []int, k int) int {
	ans := 0
	n := len(nums)
	for i := range nums {
		g := 0
		for j := i; j < n; j++ {
			g = gcd(nums[j], g)
			if g%k != 0 {
				break
			}
			if g == k {
				ans++
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
