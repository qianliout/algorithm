package main

func main() {

}

func checkPalindromeFormation(a string, b string) bool {
	return check(a, b) || check(b, a)
}

func check(a, b string) bool {
	i, j := 0, len(a)-1
	for i < j && a[i] == b[j] {
		i++
		j--
	}
	return check2(a[i:j+1]) || check2(b[i:j+1])
}

// 验证是不是回文
func check2(a string) bool {
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 参考：https://leetcode.cn/problems/split-two-strings-to-make-palindrome/solutions/1/mei-xiang-ming-bai-yi-zhang-tu-miao-dong-imvy/
