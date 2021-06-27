package main

type Rectangle struct {
	width  int
	height int
}

// 매개변수로 구조체 포인터를 받는다
func RectangleScaleA(rect *Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

// 구조체 인스턴스 형태로 받을경우 값이 모두 복사되어 원래 값의 영향을 주지않음
func RectangleScaleB(rect Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

// Go 언어에는 클래스가 없는 대신 구조체에 메서드를 연결할 수 있다.
// 리비서 변수를 사용하여 구조체 값에 접근이 가능하다.
// 함수에 구조체 형태의 매개변수를 넘기는 방식과 마찬가지로 리시버 변수를 받는 방법도 포인터와 일반 구조체 방식이 있다.
// 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	// 포인터이므로 원래 값이 변경된다.
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	// 값이 복사되었으므로 원래의 값에는 영향을 미치지 않는다.
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
