package main

import "fmt"

// https://velog.io/@proshy/JS%EC%88%9C%EC%97%B4%EC%A1%B0%ED%95%A9%EC%A4%91%EB%B3%B5%EC%88%9C%EC%97%B4-%EA%B5%AC%ED%95%98%EA%B8%B0
func main() {
	fmt.Println("순열 (3개의 카드중 2개의 카드를 선택하는 경우의 수)")

	fmt.Println(permutation([]int{1, 2, 3, 4, 5}, 3))
}

func permutation(arr []int, selectNum int) [][]int {
	/*
		배열에서 3개를 선택하는 경우를 예를 들어보면
		1. 하나의 수를 선택한다.
		2. 하나의 수를 선택했으니 남은 배열에서 2개를 선택한다. 이를 반복한다.
	*/
	var ret = [][]int{}

	/*
		1. selectNum가 1이라면 배열의 모든 원소를 배열에 담아서 return 한다.
		2. 입력받은 배열을 순회하면서 숫자를 하나씩 선택한다. = fixer
		3. fixer를 제외한 나머지 숫자들을 restArr에 담는다.
		4. permutation(restArr, selectNum - 1)를 호출한다.
		5. 그 결과에는 restArr로 selectNum - 1개의 숫자를 선택한 순열이 들어있다.
		6. combineFixer에 fixer와 permutation(restArr, selectNum - 1)의 결과를 합친 순열을 만든다.
		7. 결과를 ret에 담는다.
	*/
	if selectNum == 1 {
		for _, v := range arr {
			ret = append(ret, []int{v})
		}
		return ret
	}

	for i, v := range arr {
		fixer := v

		restArr := []int{}
		for _, w := range arr {
			if arr[i] != w {
				restArr = append(restArr, w)
			}
		}

		permutationArr := permutation(restArr, selectNum-1)
		combineFixer := [][]int{}
		for j := 0; j < len(permutationArr); j++ {
			combineFixer = append(combineFixer, append([]int{fixer}, permutationArr[j]...))
		}
		ret = append(ret, combineFixer...)
	}

	return ret
}
