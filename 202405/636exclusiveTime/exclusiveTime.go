package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	logs2 := make([]pair, len(logs))
	for i, ch := range logs {
		logs2[i] = parse(ch)
	}
	stark := make([]int, 0)
	cur := 0
	for _, ch := range logs2 {
		if ch.ex == "start" {
			if len(stark) > 0 {
				pop := stark[len(stark)-1]
				ans[pop] += ch.ti - cur
			}
			stark = append(stark, ch.id)
			cur = ch.ti
		} else {
			pop := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			ans[pop] += ch.ti - cur + 1
			cur = ch.ti + 1
		}
	}
	return ans
}

type pair struct {
	id int
	ti int
	ex string
}

func parse(ss string) pair {
	// "0:start:0"
	split := strings.Split(ss, ":")
	id, _ := strconv.Atoi(split[0])
	en, _ := strconv.Atoi(split[2])
	p := pair{
		id: id,
		ti: en,
		ex: split[1],
	}
	return p
}
