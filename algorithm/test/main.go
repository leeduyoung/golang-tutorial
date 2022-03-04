package main

import "fmt"

func main() {
	// res := solution([]int{2, 1, 3, 2}, 2)
	res := solution([]int{1, 1, 9, 1, 1, 1}, 0)
	fmt.Println(res)
}

func solution(priorities []int, location int) int {

	/*
		1. location에 해당하는 숫자를 찾는다.
		2. priorities 큐에서 location에 해당하는 숫자가 없을때까지 반복한다.
		3. 첫번째 숫자보다 큰 숫자가 큐에 존재하면 큐의 마지막에 넣는다.
		4. 없다면 큐에서 제거한다.
		5. 큐에서 제거할때 retCount 변수값(초기값 0)을 1 증가시킨다.
		6. 큐에서 제거할때 해당 값이 내가 찾던 숫자라면 반복을 제거한다.
		7. retCount 값을 리턴한다.
	*/

	var targetLocation = location
	var retCount = 0

	for len(priorities) > 0 {
		first := priorities[0]
		exists := false

		for i := 1; i < len(priorities); i++ {
			if first < priorities[i] {
				priorities = append(priorities[1:], first)
				exists = true

				if targetLocation == 0 {
					targetLocation = len(priorities) - 1
				} else {
					targetLocation--
				}

				break
			}
		}

		if !exists {
			priorities = priorities[1:]
			retCount++

			if targetLocation == 0 {
				break
			} else {
				targetLocation--
			}
		}
	}

	return retCount
}
