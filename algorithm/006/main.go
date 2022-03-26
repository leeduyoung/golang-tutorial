package main

import (
	"fmt"
	"math"
)

/*
소수 만들기
문제 설명
주어진 숫자 중 3개의 수를 더했을 때 소수가 되는 경우의 개수를 구하려고 합니다. 숫자들이 들어있는 배열 nums가 매개변수로 주어질 때, nums에 있는 숫자들 중 서로 다른 3개를 골라 더했을 때 소수가 되는 경우의 개수를 return 하도록 solution 함수를 완성해주세요.

제한사항
nums에 들어있는 숫자의 개수는 3개 이상 50개 이하입니다.
nums의 각 원소는 1 이상 1,000 이하의 자연수이며, 중복된 숫자가 들어있지 않습니다.
입출력 예
nums	result
[1,2,3,4]	1
[1,2,7,6,4]	4
입출력 예 설명
입출력 예 #1
[1,2,4]를 이용해서 7을 만들 수 있습니다.

입출력 예 #2
[1,2,4]를 이용해서 7을 만들 수 있습니다.
[1,4,6]을 이용해서 11을 만들 수 있습니다.
[2,4,7]을 이용해서 13을 만들 수 있습니다.
[4,6,7]을 이용해서 17을 만들 수 있습니다.
*/
func main() {
	res := solution([]int{1, 2, 3, 4})
	fmt.Println(res) // 1

	res = solution([]int{1, 2, 7, 6, 4})
	fmt.Println(res) // 4
}

func solution(nums []int) int {
	answer := 0

	/*
		1. 전달받은 nums 배열에서 세개를 뽑은 모든 경우의 수를 구한다.
		2. 모든 경우의수를 돌면서 소수인지 판별한다.
		3. 리턴한다.
	*/
	selectNum := 3
	res := permutation(nums, selectNum)

	for _, v := range res {
		sum := 0
		for _, w := range v {
			sum += w
		}

		if isPrime(sum) {
			answer++
		}
	}

	return answer
}

func permutation(nums []int, selectNum int) [][]int {
	/*
		순열은 n개중 selectNum 개수만큼 뽑아내는 경우의 수를 의미한다.
		nums에서 selectNum 개수만큼 뽑아내는 경우의 수를 구해보자.

		1. selectNum이 1일 경우엔 모든 수를 배열에 담아서 리턴한다.
		2. 배열을 하나씩 순회하면서 조합을 만들어나간다. fixer
		3. 선택한 숫자를 제외하고 나머지를 배열에 담는다. restArr
		4. restArr와 selectNum - 1을 사용해서 다시 재귀적으로 permutation을 호출한다.
		5. combineFixer에 재귀로 전달받은 permutation 값을 더한다.
	*/
	ret := [][]int{}

	if selectNum == 1 {
		for _, v := range nums {
			ret = append(ret, []int{v})
		}
		return ret
	}

	for i := 0; i < len(nums); i++ {
		fixer := nums[i]

		restArr := []int{}
		for _, v := range nums {
			if v != fixer {
				restArr = append(restArr, v)
			}
		}

		permutationRes := permutation(restArr, selectNum-1)
		combineFixer := [][]int{}
		for _, vv := range permutationRes {
			list := append([]int{fixer}, vv...)
			combineFixer = append(combineFixer, list)
		}

		ret = append(ret, combineFixer...)
	}

	return ret
}

func isPrime(num int) bool {
	/*
		소수는 1과 자기자신으로만 나누어지는 수
		만약 36을 예로들어보자면 루트 36까지 나누어지는수가 있는지만 확인해보면 된다.
	*/
	a := math.Sqrt(float64(num))

	for i := 2; i <= int(a); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
