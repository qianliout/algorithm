package main

func main() {

}

func trainingPlan2(actions []int) []int {
	n := len(actions)
	ans := make([]int, n)
	left, right := 0, n-1
	for _, ch := range actions {
		if ch&1 == 1 {
			ans[left] = ch
			left++
		}
	}
	for i := n - 1; i >= 0; i-- {
		if actions[i]&1 == 0 {
			ans[right] = actions[i]
			right--
		}
	}
	return ans
}

// 这个方法不能保证原顺序
func trainingPlan(actions []int) []int {
	// 定义左右指针
	left, right := 0, len(actions)-1
	for left < right {
		// 左指针向右移动，直到找到偶数
		for left < right && actions[left]%2 == 1 {
			left++
		}
		// 右指针向左移动，直到找到奇数
		for left < right && actions[right]%2 == 0 {
			right--
		}
		// 交换左右指针所指向的元素
		if left < right {
			actions[left], actions[right] = actions[right], actions[left]
		}
	}
	return actions
}
