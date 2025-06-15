package main

func main() {

}

func dailyTemperatures2(temperatures []int) []int {
	// 从右到左，
	stack := make([]int, 0) // 递减
	n := len(temperatures)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	// 从左到右
	n := len(temperatures)
	st := make([]int, 0)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		t := temperatures[i]
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}

func dailyTemperatures3(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	st := make([]int, 0)
	for i, c := range temperatures {
		// 两种解法下的判断不同，一种是> 一种是>=
		for len(st) > 0 && c > temperatures[st[len(st)-1]] {
			last := st[len(st)-1]
			ans[last] = i - last
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

func dailyTemperatures4(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	st := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		c := temperatures[i]
		for len(st) > 0 && c >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}
