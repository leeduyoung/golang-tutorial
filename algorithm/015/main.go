/*
하샤드 수
문제 설명
양의 정수 x가 하샤드 수이려면 x의 자릿수의 합으로 x가 나누어져야 합니다. 예를 들어 18의 자릿수 합은 1+8=9이고, 18은 9로 나누어 떨어지므로 18은 하샤드 수입니다. 자연수 x를 입력받아 x가 하샤드 수인지 아닌지 검사하는 함수, solution을 완성해주세요.

제한 조건
x는 1 이상, 10000 이하인 정수입니다.
입출력 예
arr	return
10	true
12	true
11	false
13	false
입출력 예 설명
입출력 예 #1
10의 모든 자릿수의 합은 1입니다. 10은 1로 나누어 떨어지므로 10은 하샤드 수입니다.

입출력 예 #2
12의 모든 자릿수의 합은 3입니다. 12는 3으로 나누어 떨어지므로 12는 하샤드 수입니다.

입출력 예 #3
11의 모든 자릿수의 합은 2입니다. 11은 2로 나누어 떨어지지 않으므로 11는 하샤드 수가 아닙니다.

입출력 예 #4
13의 모든 자릿수의 합은 4입니다. 13은 4로 나누어 떨어지지 않으므로 13은 하샤드 수가 아닙니다.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(18)) // true
	fmt.Println(solution(10)) // true
	fmt.Println(solution(12)) // true
	fmt.Println(solution(11)) // false
	fmt.Println(solution(13)) // false
}

func solution(x int) bool {
	if x == 1 {
		return true
	}

	/*
		1. 입력받은 x의 각 숫자 자리수를 순회하면서 더한다.
		2. 더한값 sum을 x로 나눈다.
		3. 이때 나머지값이 0일 경우 true
		4. 0이 아닐경우 false

		입력받은 x의 각 숫자 자리수를 어떻게 순회할 수 있을까?
		1. x를 문자열로 변환한다.
		2. 문자열을 하나씩 순회하면서 숫자로 다시 변환하고 그 값을 더한다.
	*/
	var strX string = strconv.Itoa(x)
	var sum int = 0
	for _, v := range strX {
		intChar, _ := strconv.Atoi(string(v))
		sum += intChar
	}

	return x%sum == 0
}
