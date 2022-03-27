package main

import "fmt"

func main() {
	res := permutation([]int{1, 2, 3, 4}, 3)
	fmt.Println("경우의 수: ", res)
	fmt.Println("개수: ", len(res)) // 24가지
}

/*
permutation 순열
	순열이란 서로 다른 n개중 r개를 골라 순서를 고려해 나열한 경우의 수

	예)
	중국집 메뉴 5개 중 2개의 메뉴를 순서대로 먹는 경우의 수 (순서 상관 있음)
*/
func permutation(arr []int, selectNum int) [][]int {
	ret := [][]int{}

	// 만약, selectNum이 1일 경우 arr 배열의 모든 원소를 배열로 리턴한다.
	if selectNum == 1 {
		for _, v := range arr {
			ret = append(ret, []int{v})
		}
		return ret
	}

	for _, v := range arr {
		fixer := v
		restArr := []int{}
		for _, vv := range arr {
			if v != vv {
				restArr = append(restArr, vv)
			}
		}

		result := permutation(restArr, selectNum-1)
		for _, vvv := range result {
			ret = append(ret, append([]int{fixer}, vvv...))
		}
	}

	return ret
}
