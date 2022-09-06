package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(question("Mr John Smith", 13))
}

func question(a string, size int) string {
	letters := []string{}

	for _, v := range a {
		if string(v) == " " {
			letters = append(letters, "%20")
			continue
		}
		letters = append(letters, string(v))
	}

	return strings.Join(letters, "")
}
