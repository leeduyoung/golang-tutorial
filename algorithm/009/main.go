package main

import (
	"fmt"
	"sort"
)

/*
섬 연결하기 (탐욕법)
문제 설명
n개의 섬 사이에 다리를 건설하는 비용(costs)이 주어질 때, 최소의 비용으로 모든 섬이 서로 통행 가능하도록 만들 때 필요한 최소 비용을 return 하도록 solution을 완성하세요.

다리를 여러 번 건너더라도, 도달할 수만 있으면 통행 가능하다고 봅니다. 예를 들어 A 섬과 B 섬 사이에 다리가 있고, B 섬과 C 섬 사이에 다리가 있으면 A 섬과 C 섬은 서로 통행 가능합니다.

제한사항

섬의 개수 n은 1 이상 100 이하입니다.
costs의 길이는 ((n-1) * n) / 2이하입니다.
임의의 i에 대해, costs[i][0] 와 costs[i] [1]에는 다리가 연결되는 두 섬의 번호가 들어있고, costs[i] [2]에는 이 두 섬을 연결하는 다리를 건설할 때 드는 비용입니다.
같은 연결은 두 번 주어지지 않습니다. 또한 순서가 바뀌더라도 같은 연결로 봅니다. 즉 0과 1 사이를 연결하는 비용이 주어졌을 때, 1과 0의 비용이 주어지지 않습니다.
모든 섬 사이의 다리 건설 비용이 주어지지 않습니다. 이 경우, 두 섬 사이의 건설이 불가능한 것으로 봅니다.
연결할 수 없는 섬은 주어지지 않습니다.
입출력 예

n	costs	return
4	[[0,1,1],[0,2,2],[1,2,5],[1,3,1],[2,3,8]]	4
입출력 예 설명

costs를 그림으로 표현하면 다음과 같으며, 이때 초록색 경로로 연결하는 것이 가장 적은 비용으로 모두를 통행할 수 있도록 만드는 방법입니다.

image.png
*/

func main() {
	fmt.Println("started..")
	// fmt.Println(solution(7, [][]int{{2, 3, 7}, {3, 6, 13}, {3, 5, 23}, {5, 6, 25}, {0, 1, 29}, {1, 5, 34}, {1, 2, 35}, {4, 5, 53}, {0, 4, 75}}))
	fmt.Println(solution(4, [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}}))                                   // 4
	fmt.Println(solution(5, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 3}, {3, 1, 2}, {3, 0, 4}, {2, 4, 6}, {4, 0, 7}}))             // 15
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {3, 4, 1}, {1, 2, 2}, {2, 3, 4}}))                                              // 8
	fmt.Println(solution(4, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 3}, {1, 3, 2}, {0, 3, 4}}))                                   // 9
	fmt.Println(solution(6, [][]int{{0, 1, 5}, {0, 3, 2}, {0, 4, 3}, {1, 4, 1}, {3, 4, 10}, {1, 2, 2}, {2, 5, 3}, {4, 5, 4}})) // 11
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {2, 3, 1}, {3, 4, 2}, {1, 2, 2}, {0, 4, 100}}))                                 // 6
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}, {0, 4, 4}, {1, 3, 1}}))                                   // 8

	// fail
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {0, 4, 5}, {2, 4, 1}, {2, 3, 1}, {3, 4, 1}}))   // 8
	fmt.Println(solution(5, [][]int{{0, 1, 1}, {3, 1, 1}, {0, 2, 2}, {0, 3, 2}, {0, 4, 100}})) // 104

	parent := []int{}
	for i := 0; i <= 4; i++ {
		parent = append(parent, i)
	}

	fmt.Println(parent)
	unionParent(parent, 0, 1)
	unionParent(parent, 1, 2)
	unionParent(parent, 3, 4)
	// fmt.Println(findParent(parent))
}

func solution(n int, costs [][]int) int {

	// 1. 도로를 건설하는데 필요한 비용을 기준으로 오름차순 정렬한다.
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	// 2. 각 정점이 포함된 그래프가 어디인지 저장 (초기화)
	parent := []int{}
	for i := 0; i < n; i++ {
		parent = append(parent, i)
	}

	// 3. 사이클이 발생하지 않을 경우 그래프에 포함
	sum := 0

	for i := 0; i < len(costs); i++ {
		if !findParent(parent, costs[i][0], costs[i][1]) {
			sum += costs[i][2]
			unionParent(parent, costs[i][0], costs[i][1])
		}
	}

	return sum
}

func solution2(n int, costs [][]int) int {
	min := 101

	if n == 1 {
		return min
	}

	if n == 2 {
		return costs[0][2]
	}

	/*
		1. costs에서 n-1개를 뽑는 경우의 수를 모두 구한다.
		2. 이때, 0부터 n-1까지 포함할 수 있는 경우의 수만 골라낸다.
		3. 여기서 최소값을 구한다.
	*/
	response := combination(costs, n-1)

	for _, v := range response {
		if isValidate(n, v) {
			sum := 0
			for _, w := range v {
				sum += w[2]
			}

			if sum < min {
				min = sum
			}
		}
	}

	return min
}

func isValidate(n int, data [][]int) bool {
	m := map[int]bool{}

	for _, v := range data {
		m[v[0]] = true
		m[v[1]] = true
	}

	return len(m) == n
}

func combination(arr [][]int, selectNum int) [][][]int {
	ret := [][][]int{}

	if selectNum == 1 {
		for _, v := range arr {
			ret = append(ret, [][]int{v})
		}
		return ret
	}

	for i, v := range arr {
		fixer := v
		restArr := arr[i+1:]
		result := combination(restArr, selectNum-1)

		for _, v2 := range result {
			ret = append(ret, append([][]int{fixer}, v2...))
		}
	}

	return ret
}

// 부모노드 찾기
func getParent(parent []int, x int) int {
	if parent[x] == x {
		return x
	}

	return getParent(parent, parent[x])
}

// 두 부모노드를 합치는 함수
func unionParent(parent []int, a int, b int) {
	aParent := getParent(parent, a)
	bParent := getParent(parent, b)

	if aParent < bParent {
		parent[b] = aParent
	} else {
		parent[a] = bParent
	}
}

// 같은 부모를 가지는지 확인 (즉, 두개의 노드가 동일한 그래프에 포함되어있는지 확인)
func findParent(parent []int, a int, b int) bool {
	aParent := getParent(parent, a)
	bParent := getParent(parent, b)

	return aParent == bParent
}
