package main

func main() {

}

// 方法1，从右到左
func dailyTemperatures1(temperatures []int) []int {
	t := temperatures
	st := make([]int, 0)
	n := len(t)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		// 单调,这里到底是 >= 还 >,是<= 还是 < 这个是难点
		// 如果不保留重复元素，那一定是<=或>=
		for len(st) > 0 && t[st[len(st)-1]] <= t[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	t := temperatures
	// st 中保存的是还没有找到答案的温度
	st := make([]int, 0)
	ans := make([]int, len(t))
	for i, ch := range t {
		for len(st) > 0 && t[st[len(st)-1]] < ch {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
