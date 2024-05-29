package main

import (
	"strings"
)

func main() {

}

func checkRecord(s string) bool {
	if strings.Count(s, "A") >= 2 {
		return false
	}
	start := 0

	for start < len(s) {
		if s[start] != 'L' {
			start++
			continue
		}
		end := start
		for end < len(s) && s[end] == 'L' {
			end++
		}
		if end-start >= 3 {
			return false
		}
		start = end
	}
	return true
}

func checkRecord2(s string) bool {
	/*
		A'：Absent，缺勤
		'L'：Late，迟到
		'P'：Present，到场
	*/
	if strings.Count(s, "A") >= 2 {
		return false
	}
	start := 0
	for start < len(s) {
		if s[start] == 'L' {
			end := start
			for end < len(s) && s[end] == 'L' {
				end++
			}
			if end-start >= 3 {
				return false
			}
			start = end
			continue
		}
		start++
	}
	return true
}
