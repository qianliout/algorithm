package main

func main() {

}

/*
给你一个字符串 s，模拟每秒钟的事件 i：

	如果 s[i] == 'E'，表示有一位顾客进入候诊室并占用一把椅子。
	如果 s[i] == 'L'，表示有一位顾客离开候诊室，从而释放一把椅子。

返回保证每位进入候诊室的顾客都能有椅子坐的 最少 椅子数，假设候诊室最初是 空的 。
*/
func minimumChairs1(s string) int {
	ans := 0
	stark := make([]int, 0)
	for _, ch := range s {
		if ch == 'E' {
			stark = append(stark, 1)
		} else {
			stark = stark[:len(stark)-1]
		}
		ans = max(ans, len(stark))
	}
	return ans
}

func minimumChairs(s string) int {
	ans := 0
	cnt := 0
	for _, ch := range s {
		if ch == 'E' {
			cnt++
		} else {
			cnt--
		}
		ans = max(ans, cnt)
	}
	return ans
}
