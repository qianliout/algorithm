package main

import (
	"strings"
)

func main() {

}

func largestWordCount(messages []string, senders []string) string {
	cnt := make(map[string]int)
	ans := ""

	for i, ch := range messages {
		split := strings.Split(ch, " ")
		sed := senders[i]
		cnt[sed] += len(split)

		if ans == "" || cnt[sed] > cnt[ans] || (cnt[sed] == cnt[ans] && sed < ans) {
			ans = sed
		}
	}
	return ans
}
