package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath("rzzmf\nv\n\tix\n\t\tiklav\n\t\t\ttqse\n\t\t\t\ttppzf\n\t\t\t\t\tzav\n\t\t\t\t\t\tkktei\n\t\t\t\t\t\t\thhmav\n\t\t\t\t\t\t\t\tbzvwf.txt"))
}

func lengthLongestPath(input string) int {
	res := 0
	depLenM := make(map[int]int)
	ss := strings.Split(input, "\n")
	for _, line := range ss {
		dep := strings.Count(line, "\t")
		// depLenM[dep] = max(depLenM[dep], depLenM[dep-1]+len(line)-dep) // 要减去『\t』
		depLenM[dep] = depLenM[dep-1] + len(line) - dep // 要减去『\t』
		if strings.Contains(line, ".") {
			res = max(res, depLenM[dep]+dep)
		}
	}
	return res
}
