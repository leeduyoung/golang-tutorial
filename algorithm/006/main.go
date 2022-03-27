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
		1. nums 배열에서 3개를 뽑는 경우의 수를 모두 구한다. (조합)
		2. 조합이 담긴 배열에서 모두 합을 구한다.
		3. 소수인지 판별한다.
		4. 소수라면 answer++
	*/
	res := combination(nums, 3)
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

func combination(arr []int, selectNum int) [][]int {
	ret := [][]int{}

	if selectNum == 1 {
		for _, v := range arr {
			ret = append(ret, []int{v})
		}
		return ret
	}

	for i, v := range arr {
		fixer := v
		restArr := arr[i+1:]
		result := combination(restArr, selectNum-1)

		for _, vv := range result {
			list := append([]int{fixer}, vv...)
			ret = append(ret, list)
		}
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
