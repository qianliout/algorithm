package main

import (
	"sort"
)

func main() {

}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	pairs := make([]pair, 0)
	for _, ch := range boxTypes {
		pairs = append(pairs, pair{number: ch[0], unit: ch[1]})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].unit >= pairs[j].unit })
	ans := 0
	for i := 0; i < len(pairs); i++ {
		nu := min(truckSize, pairs[i].number)
		ans += nu * pairs[i].unit
		truckSize -= nu
		if truckSize <= 0 {
			break
		}
	}
	return ans
}

type pair struct {
	number, unit int
}
