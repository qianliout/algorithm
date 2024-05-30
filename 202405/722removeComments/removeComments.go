package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeComments([]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}))
}

func removeComments(source []string) []string {
	t := make([]byte, 0)
	blockC := false

	ans := make([]string, 0)
	for _, str := range source {
		m := len(str)
		for i := 0; i < m; i++ {
			if blockC {
				if i+1 < m && str[i] == '*' && str[i+1] == '/' {
					blockC = false
					i++
				}
			} else {
				if i+1 < m && str[i] == '/' && str[i+1] == '*' {
					blockC = true
					i++
				} else if i+1 < m && str[i] == '/' && str[i+1] == '/' {
					break
				} else {
					t = append(t, str[i])
				}
			}
		}

		if !blockC && len(t) > 0 {
			ans = append(ans, string(t))
			t = make([]byte, 0)
		}
	}
	return ans
}
