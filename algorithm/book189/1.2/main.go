package main

import "fmt"

func main() {
	fmt.Println(question("god", "dog"))
}

func question(a string, b string) bool {
	fmt.Println(len(a))

	if len(a) != len(b) {
		return false
	}

	letters := [128]int{}

	for _, v := range a {
		letters[v]++
	}

	for _, v := range b {
		if letters[v] <= 0 {
			return false
		}

		letters[v]--
	}

	return true
}
