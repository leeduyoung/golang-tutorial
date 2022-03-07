package main

import (
	"fmt"
)

func main() {
	res := solution(13, 17)
	fmt.Println(res)
}

func solution(left int, right int) int {

	/*
		1. left부터 right까지 순회하면서
		2. 약수의 갯수를 구하는 함수를 실행시키고
		3. 홀수면 - 짝수면 +
	*/

	var ret int = 0

	for i := left; i <= right; i++ {
		c := divisorCount(i)
		if c%2 == 0 {
			ret += i
		} else {
			ret -= i
		}
	}

	return ret
}

func divisorCount(n int) int {
	var count = 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}
