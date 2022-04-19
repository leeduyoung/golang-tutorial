package main

import (
	"fmt"
	"strings"
)

/*
큰 수 만들기 (탐욕법)
문제 설명
어떤 숫자에서 k개의 수를 제거했을 때 얻을 수 있는 가장 큰 숫자를 구하려 합니다.

예를 들어, 숫자 1924에서 수 두 개를 제거하면 [19, 12, 14, 92, 94, 24] 를 만들 수 있습니다. 이 중 가장 큰 숫자는 94 입니다.

문자열 형식으로 숫자 number와 제거할 수의 개수 k가 solution 함수의 매개변수로 주어집니다. number에서 k 개의 수를 제거했을 때 만들 수 있는 수 중 가장 큰 숫자를 문자열 형태로 return 하도록 solution 함수를 완성하세요.

제한 조건
number는 2자리 이상, 1,000,000자리 이하인 숫자입니다.
k는 1 이상 number의 자릿수 미만인 자연수입니다.
입출력 예
number	k	return
"1924"	2	"94"
"1231234"	3	"3234"
"4177252841"	4	"775841"

*/

func main() {
	// fmt.Println(solution("1924", 2))       // 94
	// fmt.Println(solution("1231234", 3))    // 3234
	fmt.Println(solution("4177252841", 4)) // 775841
}

func solution(number string, k int) string {

	/*
		1. number의 길이 = n 이라고 할떄, n - k 길이의 빈배열(arr)을 하나 만든다.
		2. number를 하나씩 순회하면서
		3. arr에 값을 채운다.
		4. 배열이 비여있는 최초에는 첫번째 값을 넣는다.
		5. 두번째 값이 arr의 첫번째 값보다 크고 && 뒤에 (남은 숫자 - 1)이 arr의 비어있는 수보다 많으면 값을 변경
		6. 나머지는 모두 다음 위치에 넣기
	*/

	n := len(number)
	arr := []string{}
	for i := 0; i < n-k; i++ {
		arr = append(arr, "0")
	}

	arrIndex := 0
	for i, v := range number {
		value := fmt.Sprintf("%c", v)

		if i == 0 {
			arr[arrIndex] = value
			arrIndex++
			continue
		}

		if arr[arrIndex-1] < value && (n-i-1) >= (n-k-arrIndex) {
			arr[arrIndex-1] = value
			continue
		}

		arr[arrIndex] = value
		arrIndex++
	}

	return strings.Join(arr, "")
}
