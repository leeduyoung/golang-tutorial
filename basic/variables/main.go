package main

import "fmt"

func main() {
	// 상수값을 0부터 순차적으로 부여하기 위한 팁
	const (
		Apple = iota
		Microsoft
		Google
	)

	fmt.Println("Apple: ", Apple)
	fmt.Println("Microsoft: ", Microsoft)
	fmt.Println("Google: ", Google)

	// TYPE CONVERSION
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(u)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

}
