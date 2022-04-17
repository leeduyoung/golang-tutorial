package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution(4, [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}}))                                   // 4
	fmt.Println(solution(5, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 3}, {3, 1, 2}, {3, 0, 4}, {2, 4, 6}, {4, 0, 7}}))             // 15
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {3, 4, 1}, {1, 2, 2}, {2, 3, 4}}))                                              // 8
	fmt.Println(solution(4, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 3}, {1, 3, 2}, {0, 3, 4}}))                                   // 9
	fmt.Println(solution(6, [][]int{{0, 1, 5}, {0, 3, 2}, {0, 4, 3}, {1, 4, 1}, {3, 4, 10}, {1, 2, 2}, {2, 5, 3}, {4, 5, 4}})) // 11
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {2, 3, 1}, {3, 4, 2}, {1, 2, 2}, {0, 4, 100}}))                                 // 6
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}, {0, 4, 4}, {1, 3, 1}}))                                   // 8
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {0, 4, 5}, {2, 4, 1}, {2, 3, 1}, {3, 4, 1}}))                                   // 8
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {3, 1, 1}, {0, 2, 2}, {0, 3, 2}, {0, 4, 100}}))                                 // 104
}

func solution(n int, costs [][]int) int {

	/*
		1. 최소비용만 골라내면 되기 때문에, 도로를 건설하는데 드는 비용을 오름차순으로 정렬한다.
		2. 각 정점이 포함된 그래프가 어디인지 저장 -> 각 노드의 부모 노드 값을 저장하는 배열을 만든다.
		3. 사이클이 발생하지 않을 경우는 그래프에 포함
	*/
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	parent := []int{}
	for i := 0; i < n; i++ {
		parent = append(parent, i)
	}

	sum := 0

	for i := 0; i < len(costs); i++ {
		if !findParent(parent, costs[i][0], costs[i][1]) {
			sum += costs[i][2]
			unionParent(parent, costs[i][0], costs[i][1])
		}
	}

	return sum
}

// 부모노드 찾기
func getParent(parent []int, x int) int {
	if parent[x] == x {
		return x
	}

	return getParent(parent, parent[x])
}

// 같은 부모를 가지는지 확인 (즉, 두개의 노드가 동일한 그래프에 포함되어 있는지 확인)
func findParent(parent []int, a int, b int) bool {
	aParent := getParent(parent, a)
	bParent := getParent(parent, b)

	return aParent == bParent
}

// 두 부모노드를 합치는 함수
func unionParent(parent []int, a int, b int) {
	aParent := getParent(parent, a)
	bParent := getParent(parent, b)

	if aParent < bParent {
		parent[bParent] = aParent
	} else {
		parent[aParent] = bParent
	}
}
