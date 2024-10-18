package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"xga", "xfb", "yfa"}))
}

func minDeletionSize(strs []string) int {
	cnt := 0
	n := len(strs[0])
	m := len(strs)
	for i := 0; i < n; i++ {
		need := false
		same := false
		for j := 1; j < m; j++ {
			if strs[j-1][j] == strs[j][i] {
				same = true
			}
			if strs[j-1][i] > strs[j][i] {
				need = true
				break
			}
		}
		if need {
			cnt++
		} else if !same {
			break
		}
	}
	return cnt
}
