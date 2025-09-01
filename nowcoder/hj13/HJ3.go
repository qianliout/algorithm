package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	n, _ := strconv.Atoi(input.Text())

	data := make([]int, 501)

	for i := 0; i < n; i++ {
		input.Scan()
		ch, _ := strconv.Atoi(input.Text())
		data[ch] = ch
	}
	for _, ch := range data {
		if ch > 0 {
			fmt.Println(ch)
		}
	}
}
