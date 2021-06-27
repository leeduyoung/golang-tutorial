package main

import "fmt"

func PointerTest() {
	// 포인터형 변수 선언
	var numPtr *int     // 포인터형 변수를 선언하면 nil로 초기화 된다.
	fmt.Println(numPtr) // nil

	var numPtr2 *int = new(int) // 빈 포인터형 변수는 바로 사용할 수 없으므로 new 함수로 메모리를 할당한다.
	fmt.Println(numPtr2)        // 메모리 주소. 시스템 마다, 실행할때마다 달라진다.

	// 포인터형 변수에 값을 대입하거나, 가져오려면 역참조를 사용한다.
	// 역참조로 포인터형 변수에 값을 대입
	*numPtr2 = 1
	fmt.Println(numPtr2)
	fmt.Println(*numPtr2) // 포인터형 변수의 값을 가져오기

	// 변수를 선언할때 *를 붙이면 포인터형 변수가 되지만, 변수를 사용할 때 *를 붙이면 역참조가 된다.
	// *numPtr = 1은 numPtr에 저장된 메모리 주소로 접근하여 값을 대입합니다.
	// 그리고 *numPtr은 numPtr에 저장된메모리 주소에 접근하여 값을 가져옵니다.

	// 일반 변수에 참조(레퍼런스)를 사용하면 포인터형 변수에 대입할 수 있다.
	// 변수 앞에 &를 붙이면 해당 변수의 메모리 주소를 뜻한다.
	var num int = 1
	var numPtr3 *int = &num // 참조로 num 변수의 메모리 주소를 구하여 numPtr3 포인터 변수에 대입한다.
	fmt.Println(numPtr3)
	fmt.Println(&num)
}
