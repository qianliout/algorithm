package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	ans := make([][]string, 0)
	n := len(searchWord)
	for i := 1; i <= n; i++ {
		cur := searchWord[:i]
		ans = append(ans, find(products, cur))
	}
	return ans
}

func find(products []string, word string) []string {
	le, ri := 0, len(products)
	n := len(word)
	for le < ri {
		// 左边界
		mid := le + (ri-le)/2
		// if products[mid][:n] == word {
		if products[mid] >= word {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	ans := make([]string, 0)
	for i := le; i < min(len(products), le+3); i++ {
		mx := min(n, len(products[i]))
		if products[i] >= word && products[i][:mx] == word {
			ans = append(ans, products[i])
		}
	}

	return ans
}
