package main

import (
	"fmt"
	"sort"
)

/*

N개의 최소공배수
문제 설명
두 수의 최소공배수(Least Common Multiple)란 입력된 두 수의 배수 중 공통이 되는 가장 작은 숫자를 의미합니다. 예를 들어 2와 7의 최소공배수는 14가 됩니다. 정의를 확장해서, n개의 수의 최소공배수는 n 개의 수들의 배수 중 공통이 되는 가장 작은 숫자가 됩니다. n개의 숫자를 담은 배열 arr이 입력되었을 때 이 수들의 최소공배수를 반환하는 함수, solution을 완성해 주세요.

제한 사항
arr은 길이 1이상, 15이하인 배열입니다.
arr의 원소는 100 이하인 자연수입니다.
입출력 예
arr	result
[2,6,8,14]	168
[1,2,3]	6

*/

func main() {
	fmt.Println(solution([]int{2, 6, 8, 14})) // 168
	fmt.Println(solution([]int{1, 2, 3}))     // 6

	// Failed
	fmt.Println(solution([]int{3, 4, 9, 16})) // 144
	fmt.Println(solution([]int{2, 3, 4}))     // 12
}

func solution(arr []int) int {
	var answer = 0

	// 1. arr 오름차순 정렬
	sort.Ints(arr)

	// 2. 배열에서 1은 제외
	newArr := []int{}
	for _, v := range arr {
		if v == 1 {
			continue
		}
		newArr = append(newArr, v)
	}

	// 3. 가장 작은수부터 2까지 순회
	for i := newArr[0]; i > 1; i-- {
		var bool = true
		var val = []int{}

		for _, v := range newArr {
			if v%i != 0 {
				bool = false
				break
			}
			val = append(val, v/i)
		}

		if bool {
			sum := 1
			for _, w := range val {
				sum *= w
			}
			answer = i * sum
		}
	}

	if answer == 0 {
		ret := 1
		for _, v := range newArr {
			ret *= v
		}
		answer = ret
	}

	return answer
}

// 여러수의 최소 공배수는 순서에 상관없이 두개씩 순차적으로 구하면 된다.

// 최대 공약수 - gcd (유클리드 호제법)
// TODO:

// 최소 공배수 -lcm
// TODO:
