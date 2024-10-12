package main

import "fmt"

func main() {
	fmt.Println(smallestBeautifulString("da", 8))
	fmt.Println(smallestBeautifulString("dc", 4))
}

// 这里的逻辑控制的真的是绝了,膜拜

func smallestBeautifulString(s string, k int) string {
	limit := 'a' + byte(k)
	ss := []byte(s)
	n := len(s)
	i := len(s) - 1
	ss[i]++
	for i < n && i >= 0 {
		if ss[i] >= limit {
			if i == 0 {
				return ""
			}
			ss[i] = 'a'
			i--
			ss[i]++
			continue
		}
		// 检测是否有回文串
		if (i >= 1 && ss[i] == ss[i-1]) || (i >= 2 && ss[i] == ss[i-2]) {
			// 如果前面有问题就先检测前面
			ss[i]++
			continue
		}
		i++ // 再返回回来检测后面
	}

	return string(ss)
}
