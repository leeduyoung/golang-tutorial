package main

import (
	"fmt"
)

func main() {
	fmt.Println(question("abccbagg"))
}

func question(word string) bool {
	m := make(map[string]int)
	for _, v := range word {
		str := string(v)
		if _, exists := m[str]; !exists {
			m[str] = 1
			continue
		}
		m[str] += 1
	}

	fmt.Println(m)

	oddCount := 0
	for _, v := range m {
		if v%2 == 1 {
			oddCount++
		}

		if oddCount > 1 {
			return false
		}
	}

	return true
}
