package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}))
}

func findDuplicate(paths []string) [][]string {
	mp := make(map[string][]pair)
	for _, ch := range paths {
		pairs := parse(ch)
		for i := range pairs {
			pa := pairs[i]
			if mp[pa.c] == nil {
				mp[pa.c] = make([]pair, 0)
			}
			mp[pa.c] = append(mp[pa.c], pa)
		}
	}
	ans := make([][]string, 0)
	for _, pars := range mp {
		res := make([]string, 0)
		for i := range pars {
			pa := pars[i]
			res = append(res, pa.p)
		}
		if len(res) >= 2 {
			ans = append(ans, res)
		}
	}
	return ans
}

type pair struct {
	c string
	p string
}

func parse(s string) []pair {
	ans := make([]pair, 0)
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	if len(split) == 0 || split[0] == "" {
		return ans
	}
	p := split[0]
	for i := 1; i < len(split); i++ {
		if !strings.Contains(split[i], "(") || !strings.Contains(split[i], ")") {
			continue
		}
		i2 := strings.Split(split[i], "(")
		if len(i2) != 2 {
			continue
		}
		fi := i2[0]
		c := strings.TrimRight(i2[1], ")")
		ans = append(ans, pair{
			c: c,
			p: fmt.Sprintf("%s/%s", p, fi),
		})
	}
	return ans
}
