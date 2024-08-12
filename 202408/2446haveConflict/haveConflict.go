package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(haveConflict([]string{"10:00", "11:00"}, []string{"14:00", "15:00"}))

}

func haveConflict2(event1 []string, event2 []string) bool {
	as := cal(event1[0])
	ae := cal(event1[1])
	bs := cal(event2[0])
	be := cal(event2[1])
	return !(as > be || ae < bs)
}

func cal(s string) int {
	split := strings.Split(s, ":")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	return a*60 + b
}

// 可以直接比较
func haveConflict(event1 []string, event2 []string) bool {
	as := event1[0]
	ae := event1[1]
	bs := event2[0]
	be := event2[1]
	return !(as > be || ae < bs)
}
