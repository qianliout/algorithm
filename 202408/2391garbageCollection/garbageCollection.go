package main

import (
	"strings"
)

func main() {

}

func garbageCollection(garbage []string, travel []int) int {
	ans := 0
	for _, ch := range garbage {
		ans += len(ch)
	}
	for _, ch := range travel {
		ans += ch * 3
	}
	for _, c := range []byte("MPG") {
		for j := len(garbage) - 1; j > 0; j-- {
			cnt := strings.Count(garbage[j], string(c))
			if cnt > 0 {
				break
			}
			ans -= travel[j-1]
		}
	}
	return ans
}
