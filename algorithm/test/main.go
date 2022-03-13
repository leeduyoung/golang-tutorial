package main

import (
	"fmt"
)

func main() {
	res := solution([]int{0, 0, 0, 0, 0, 0}, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(res)
}

func solution(lottos []int, win_nums []int) []int {

	rankMap := map[int]int{
		6: 1,
		5: 2,
		4: 3,
		3: 4,
		2: 5,
		1: 6,
		0: 6,
	}

	equalsCount := 0
	zeroCount := 0

	for _, v := range lottos {
		if v == 0 {
			zeroCount++
		}

		for _, w := range win_nums {
			if v == w {
				equalsCount++
			}
		}
	}

	best := rankMap[equalsCount+zeroCount]
	worst := rankMap[equalsCount]

	return []int{best, worst}
}
