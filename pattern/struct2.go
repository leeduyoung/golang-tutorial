package main

type Rectangle struct {
	width int
	height int
}

// 매개변수로 구조체 포인터를 받는다
func RectangleScaleA(rect *Rectangle, factor int) {
	rect.width *= factor
	rect.height*= factor
}

// 구조체 인스턴스 형태로 받을경우 값이 모두 복사되어 원래 값의 영향을 주지않음 
func RectangleScaleB(rect Rectangle, factor int) {
	rect.width *= factor
	rect.height*= factor
}