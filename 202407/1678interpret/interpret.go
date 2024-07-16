package main

import (
	"strings"
)

func main() {

}

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	command = strings.ReplaceAll(command, "()", "o")
	return command
}
