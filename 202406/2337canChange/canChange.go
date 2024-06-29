package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(canChange("_L__R__RL", "L_____RLR"))
}

func canChange(start string, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}

	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == '_' {
			continue
		}
		for j < len(start) && target[j] == '_' {
			j++
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		j++
	}
	return true
}
