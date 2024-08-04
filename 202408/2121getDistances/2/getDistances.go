package main

func main() {

}

func getDistances(arr []int) []int64 {
	ans := make([]int64, len(arr))
	cnt := make(map[int][]int)

	for i, ch := range arr {
		cnt[ch] = append(cnt[ch], i)
	}
	for _, p := range cnt {
		sum := int64(0)
		for _, i := range p {
			sum += int64(i - p[0])
		}
		ans[p[0]] = sum
		n := len(p)
		for i := 1; i < n; i++ {
			// 这个方法一定得会
			// 然后我们不断向右计算下一个元素的间隔和。对比 p 中的第 i−1 个元素与第 i 个元素，观察他们间隔和的变化量：
			//    左边有 i 个元素的间隔变大了 p[i]−p[i−1]；
			//    右边有 n−i 个元素的间隔变小了 p[i]−p[i−1]。
			sum += int64(2*i-n) * int64(p[i]-p[i-1])
			ans[p[i]] = sum
		}
	}
	return ans
}

func cal(nums []int, a int) int64 {
	ans := 0
	for _, ch := range nums {
		if ch == a {
			continue
		}
		ans += abs(ch - a)
	}
	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
