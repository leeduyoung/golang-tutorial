package main

import "fmt"

/*
	덕 타이핑
	각 값이나 인스턴스의 실제 타입은 상관하지않고 구현도니 메서드로만 타입을 판단하는 방식을 덕 타이핑이라고 한다.

	오리든 사람이든 상관없이 quack(), feathers()라는 메서드만 가졌다면 모두 같은 인터페이스로 처리가 가능하다.
*/

type Duck struct{}

func (d Duck) quack() {
	fmt.Println("꽥~!")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색과 회색 털을 가지고 있다.")
}

type Person3 struct{}

func (p Person3) quack() {
	fmt.Println("사람은 오리 흉내를 낸다. 꽥~!")
}

func (p Person3) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여준다.")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}
