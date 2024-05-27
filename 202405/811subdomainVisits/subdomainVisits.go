package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func subdomainVisits(cpdomains []string) []string {
	ans := make(map[string]int)
	for i := range cpdomains {
		cacal(cpdomains[i], ans)
	}
	res := make([]string, 0)
	for k, v := range ans {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}

func cacal(dom string, ans map[string]int) {
	split1 := strings.Split(dom, " ")
	if len(split1) != 2 {
		return
	}
	cnt, _ := strconv.Atoi(split1[0])

	split := strings.Split(split1[1], ".")
	for i := 0; i < len(split); i++ {
		key := strings.Join(split[i:], ".")
		ans[key] += cnt
	}
}
