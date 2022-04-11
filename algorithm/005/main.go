/*
programmers

소수 찾기 (완전 탐색)
문제 설명
한자리 숫자가 적힌 종이 조각이 흩어져있습니다. 흩어진 종이 조각을 붙여 소수를 몇 개 만들 수 있는지 알아내려 합니다.

각 종이 조각에 적힌 숫자가 적힌 문자열 numbers가 주어졌을 때, 종이 조각으로 만들 수 있는 소수가 몇 개인지 return 하도록 solution 함수를 완성해주세요.

제한사항
numbers는 길이 1 이상 7 이하인 문자열입니다.
numbers는 0~9까지 숫자만으로 이루어져 있습니다.
"013"은 0, 1, 3 숫자가 적힌 종이 조각이 흩어져있다는 의미입니다.
입출력 예
numbers	return
"17"	3
"011"	2
입출력 예 설명
예제 #1
[1, 7]으로는 소수 [7, 17, 71]를 만들 수 있습니다.

예제 #2
[0, 1, 1]으로는 소수 [11, 101]를 만들 수 있습니다.

11과 011은 같은 숫자로 취급합니다.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// res := solution("17")
	// fmt.Println(res) // 3

	// res = solution("011")
	// fmt.Println(res) // 2

	fmt.Println(permutation([]string{"1", "2", "3"}, 3))
}

func solution(numbers string) int {
	/*
	   1. 입력값을 숫자로 변환
	   2. 주어진 숫자로 만들 수 있는 모든 수를 찾는다.
	   3. 찾아낸 수 가운데 소수를 찾는다.
	*/

	var ret = make(map[int]struct{})
	var arr []string = strings.Split(numbers, "")

	k := [][]string{}
	for i := 1; i <= len(arr); i++ {
		p := permutation(arr, i)
		k = append(k, p...)
	}

	for i := 0; i < len(k); i++ {
		val := ""
		for j := 0; j < len(k[i]); j++ {
			val += k[i][j]
		}
		rval, _ := strconv.Atoi(val)
		if isPrime(rval) {
			ret[rval] = struct{}{}
		}
	}

	return len(ret)
}

func permutation(arr []string, n int) [][]string {
	var ret [][]string

	if n == 1 {
		for i := 0; i < len(arr); i++ {
			ret = append(ret, []string{arr[i]})
		}
		return ret
	}

	for i, v := range arr {
		fixer := v
		newArr := []string{}
		for ii, vv := range arr {
			// 주의할 부분. 값으로 비교할 경우 중복된 숫자가 있는경우 문제가 발생한다.
			if ii != i {
				newArr = append(newArr, vv)
			}
		}

		permutationArr := permutation(newArr, n-1)
		combineFixer := [][]string{}
		for _, w := range permutationArr {
			combineFixer = append(combineFixer, append([]string{fixer}, w...))
		}
		ret = append(ret, combineFixer...)
	}
	return ret
}

// isPrime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
