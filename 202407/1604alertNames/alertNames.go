package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func alertNames(keyName []string, keyTime []string) []string {
	res := make([]pair, 0)
	for i := range keyName {
		res = append(res, pair{name: keyName[i], t: parse(keyTime[i])})
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i].name < res[j].name {
			return true
		} else if res[i].name > res[j].name {
			return false
		}
		return res[i].t < res[j].t
	})
	ans := make([]string, 0)
	for i := 0; i < len(res)-2; i++ {
		cur := res[i]
		nex := res[i+2]

		if cur.name == nex.name && nex.t-cur.t <= 60 {
			ans = append(ans, cur.name)
		}

	}
	ans = dup(ans)
	return ans
}

type pair struct {
	name string
	t    int
}

func parse(t string) int {
	split := strings.Split(t, ":")

	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	return a*60 + b
}

func count(a []pair) []pair {
	ans := make([]pair, 0)
	cnt := make(map[string]int)
	for _, ch := range a {
		cnt[ch.name]++
	}

	return ans
}

func dup(a []string) []string {
	exist := make(map[string]bool)
	res := make([]string, 0)
	for _, ch := range a {
		if !exist[ch] {
			res = append(res, ch)
			exist[ch] = true
		}
	}
	return res
}
