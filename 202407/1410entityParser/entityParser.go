package main

import (
	"strings"
)

func main() {

}

/*

   双引号：字符实体为 &quot; ，对应的字符是 " 。
   单引号：字符实体为 &apos; ，对应的字符是 ' 。
   与符号：字符实体为 &amp; ，对应对的字符是 & 。
   大于号：字符实体为 &gt; ，对应的字符是 > 。
   小于号：字符实体为 &lt; ，对应的字符是 < 。
   斜线号：字符实体为 &frasl; ，对应的字符是 / 。
*/

func entityParser(text string) string {
	rep := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
		// "&amp;":"&",
	}
	for k, v := range rep {
		text = strings.ReplaceAll(text, k, v)
	}
	// 一定要最后替换
	text = strings.ReplaceAll(text, "&amp;", "&")
	return text
}
