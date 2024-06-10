package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	sum1 := 0
	for i := range customers {
		if grumpy[i] == 0 {
			sum1 += customers[i]
		}
	}
	sum2 := 0

	le, ri := 0, 0
	wind := 0
	for le <= ri && ri < n {
		if grumpy[ri] == 1 {
			wind += customers[ri]
		}
		ri++
		for ri-le > minutes {
			if grumpy[le] == 1 {
				wind -= customers[le]
			}
			le++
		}
		if ri-le == minutes {
			sum2 = max(sum2, wind)
		}

	}
	return sum1 + sum2
}
