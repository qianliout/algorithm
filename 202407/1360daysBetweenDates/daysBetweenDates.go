package main

import (
	"time"
)

func main() {

}

func daysBetweenDates(date1 string, date2 string) int {
	lay := "2006-01-02"
	n1, _ := time.Parse(lay, date1)
	n2, _ := time.Parse(lay, date2)
	n := n1.Sub(n2)
	ans := int(n.Hours() / 24)
	if ans < 0 {
		return -ans
	}
	return ans
}
