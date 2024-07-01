package main

func main() {

}

func minimumRounds(tasks []int) int {
	cnt := make(map[int]int)
	for _, ch := range tasks {
		cnt[ch]++
	}
	ans := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		}

		/*
			分类讨论：
			    如果 c=1，无法完成，返回 −1。
			    如果 c=3k (k≥1)，只用「减少 3」就能完成，轮数为 c/3。
			    如果 c=3k+1 (k≥1)，即 c=3k′+4 (k′≥0)，我们可以先把 c 减少到 4，然后使用两次「减少 2」，轮数为 (c-4)/3+2 也等于(c+2)/3。
			    如果 c=3k+2 (k≥1)，我们可以先把 c 减少到 2，然后使用一次「减少 2」，轮数为 (c-2)/3+1 也等于（c+1）/ 3 。
			综上所述，对于 c (c≥2) 个相同难度级别的任务，最少需要操作 (c+2)/3 这里主要使用的除法向下取整数
		*/
		ans += (v + 2) / 3
	}
	return ans
}
