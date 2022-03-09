package main

import (
	"fmt"
	"math"
)

/*
	소수
	소수는 2,3,5,7,... 처럼 1과 자기자신만으로 나누어 떨어지는 1보다 큰 양의 정수이다.
*/
func main() {
	fmt.Println("소수 판별")

	var count int = 0
	for i := 0; i <= 100; i++ {
		if isPrime(i) {
			count++
		}
	}
	fmt.Println(count)

	count = 0
	for i := 0; i <= 100; i++ {
		if isPrime2(i) {
			count++
		}
	}
	fmt.Println(count)
}

// isPrime
// n이 1보다 클때, 2부터 n까지의 수를 차례대로 나누어 떨어지는 수가 없으면 소수이다.
func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// isPrime2
// n이 36이라고 가정하면, 36의 약수는 1, 2, 3, 4, 6, 9, 12, 18, 36이다.
// 이를 이용하면 루트 n까지만 반복해보면 이 수가 소수인지 판별할 수 있다.
func isPrime2(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
