package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

}

func reformatDate(dateStr string) string {
	d := []string{"1st", "2nd", "3rd"}

	day := make(map[string]int)
	for i, ch := range d {
		day[ch] = i + 1
	}
	for i := 4; i <= 31; i++ {
		key := fmt.Sprintf("%dth", i)
		day[key] = i
	}
	day["11st"] = 11
	day["12nd"] = 12
	day["13rd"] = 13

	day["21st"] = 21
	day["22nd"] = 22
	day["23rd"] = 23
	day["31st"] = 31

	m := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	month := make(map[string]int)
	for i, ch := range m {
		month[ch] = i + 1
	}

	split := strings.Split(dateStr, " ")
	if len(split) != 3 {
		return ""
	}
	ny, _ := strconv.Atoi(split[2])
	nd := day[split[0]]
	nm := month[split[1]]
	data := time.Date(ny, time.Month(nm), nd, 0, 0, 0, 0, time.UTC)

	return data.Format(time.DateOnly)
}
