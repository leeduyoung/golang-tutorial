package main

import "fmt"

/*
피보나치 수
문제 설명
피보나치 수는 F(0) = 0, F(1) = 1일 때, 1 이상의 n에 대하여 F(n) = F(n-1) + F(n-2) 가 적용되는 수 입니다.

예를들어

F(2) = F(0) + F(1) = 0 + 1 = 1
F(3) = F(1) + F(2) = 1 + 1 = 2
F(4) = F(2) + F(3) = 1 + 2 = 3
F(5) = F(3) + F(4) = 2 + 3 = 5
와 같이 이어집니다.

2 이상의 n이 입력되었을 때, n번째 피보나치 수를 1234567으로 나눈 나머지를 리턴하는 함수, solution을 완성해 주세요.

제한 사항
n은 2 이상 100,000 이하인 자연수입니다.
입출력 예
n	return
3	2
5	5
입출력 예 설명
피보나치수는 0번째부터 0, 1, 1, 2, 3, 5, ... 와 같이 이어집니다.
*/

func solution(n int) int {
	fibMap := map[int]int{
		0: 0,
		1: 1,
	}
	return fibonaci(n, &fibMap)
}

func fibonaci(n int, fibMap *map[int]int) int {
	fMap := *fibMap

	if val, exists := fMap[n]; exists {
		return val
	}

	// 주의사항, 마지막 결과값에 1234567를 나눈 나머지 값을 계산할 경우 시간초과 혹은 자료형 overflow가 발생할 수 있다.
	// 따라서, 매 연산에서 나온 결과값에 1234567로 나눈 나머지를 계산할 경우 이를 해결할 수 있다.
	// 즉, (A+B) % C = ((A % C) + (B % C)) % C 는 동일한 결과를 나타낸다.
	fMap[n] = (fibonaci(n-1, fibMap) + fibonaci(n-2, fibMap)) % 1234567
	return fMap[n]
}

func main() {
	fmt.Println("Input: 3, Output: ", solution(3)) // 2
	fmt.Println("Input: 5, Output: ", solution(5)) // 5
}

func fibonaciOrigin(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibonaciOrigin(n-1) + fibonaciOrigin(n-2)
}
