package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maximumBinaryString("1100"))
}

// 最多只能有一个0，且1只能右移
// 那么当上述操作完成时，有 cnt2个 1 被挤到答案的末尾，那唯一的 0 就在这 cnt2个 1 的左边：

func maximumBinaryString(binary string) string {
	idx := strings.Index(binary, "0")
	if idx < 0 { // 全是1
		return binary
	}
	cnt2 := strings.Count(binary[idx:], "1")
	n := len(binary)
	return strings.Repeat("1", n-cnt2-1) + "0" + strings.Repeat("1", cnt2)
}
