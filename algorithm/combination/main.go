package main

import "fmt"

func main() {
	res := combination([]int{1, 2, 3, 4}, 3)
	fmt.Println("경우의 수: ", res)
	fmt.Println("개수: ", len(res)) // 4가지
}

/*
combination 조합
	조합이란 서로 다른 n개중 r개를 선택하는 경우의 수

	에)
	중국집 메뉴 5개 중 2개의 메뉴를 주문하는 경우의 수 (순서 상관 없음)
*/
func combination(arr []int, selectNum int) [][]int {
	ret := [][]int{}

	// 만약, selectNum이 1일 경우 모든 원소의 배열을 리턴한다.
	if selectNum == 1 {
		for _, v := range arr {
			ret = append(ret, []int{v})
		}
		return ret
	}

	for i, v := range arr {
		fixer := v
		restArr := arr[i+1:] // 조합에서는 순서가 상관 없으므로, 이전 원소를 제외하고 시작한다.
		result := combination(restArr, selectNum-1)

		for _, v2 := range result {
			ret = append(ret, append([]int{fixer}, v2...))
		}
	}

	return ret
}
