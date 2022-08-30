package main

import (
	"fmt"
	"sort"
)

func solution(jobs [][]int) int {
	var answer = 0
	/*
		1. jobs를 순회하면서 배열의 합한 숫자가 작은 순서대로 정렬한다.
		2. 계산한다.
	*/

	sort.Slice(jobs, func(i, j int) bool {
		var (
			isum = 0
			jsum = 0
		)

		for _, v := range jobs[i] {
			isum += v
		}
		for _, v := range jobs[j] {
			jsum += v
		}

		return isum < jsum
	})
	fmt.Println("jobs: ", jobs)

	for i, v := range jobs {
		if i == 0 {
			answer += v[1]
		}
	}

	return answer
}

func main() {
	fmt.Println(solution([][]int{{0, 3}, {1, 9}, {2, 6}}) == 9) // 9
}
