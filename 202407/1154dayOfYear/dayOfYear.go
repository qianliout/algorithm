package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}

func dayOfYear(date string) int {
	ly := "2006-01-02"
	parse, _ := time.Parse(ly, date)

	start := time.Date(parse.Year(), 1, 0, 0, 0, 0, 0, time.UTC)
	return int(parse.Sub(start).Hours()) / 24
}
