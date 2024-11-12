package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))
}

func simplifyPath(path string) string {
	ans := make([]string, 0)
	split := strings.Split(path, "/")
	for _, ch := range split {
		ch = strings.TrimSpace(ch)
		if ch == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
			continue
		} else if ch == "/" || ch == "." || ch == "" {
			continue
		}
		ans = append(ans, ch)
	}
	s := strings.Join(ans, "/")
	s = "/" + s
	return s
}

// 在 Unix 风格的文件系统中规则如下：
// 一个点 '.' 表示当前目录本身。
// 此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
// 任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
// 任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
// 返回的 简化路径 必须遵循下述格式：
// 始终以斜杠 '/' 开头。
// 两个目录名之间必须只有一个斜杠 '/' 。
// 最后一个目录名（如果存在）不能 以 '/' 结尾。
// 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
