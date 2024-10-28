package main

import (
	"fmt"
	"strings"
)

func main() {

}

func numUniqueEmails(emails []string) int {
	cnt := make(map[Email]int)
	for _, ch := range emails {
		e, err := getEmail(ch)
		if err == nil {
			cnt[e]++
		}
	}
	return len(cnt)
}

type Email struct {
	First  string
	Second string
}

func getEmail(str string) (Email, error) {
	split := strings.Split(str, "@")
	if len(split) != 2 {
		return Email{}, fmt.Errorf("error")
	}

	first := split[0]
	idx := strings.Index(first, "+")
	if idx != -1 {
		first = first[:idx]
	}
	first = strings.ReplaceAll(first, ".", "")
	return Email{First: first, Second: split[1]}, nil
}
