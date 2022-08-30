package main

import "fmt"

/*
중복이 없는가
문자열이 주어졌을때 이 문자열에 같은 문자가 중복되어 등장하는지 확인하는 알고리즘을 작성하라.
자료구조를 추가로 사용하지 않고 풀 수 있는 알고리즘 또한 고민하라.

사전지식
	- 문자의 집합인 문자열을 순회할 수 있는가?
	- 한글과 영문 모두 문자열 순회할 수 있는가? (index 순회를 할 경우 구분 필요)
	- rune 타입에 대하여 알고있는가? (rune은 유니코드를 표현하는 타입으로 int32 타입의 별칭과 같다)
	- 유니코드 인코딩에서 한글은 3byte를 사용하고, 영어는 1byte를 사용한다.
	- 유니코드와 아스키코드의 개념

좋은 질문사항
	- input으로 들어갈 수 있는 문자가 아스키코드 인지 또는 영어 알파벳인지에 대해 자세하게 물어보면 좋다.
	- 만약, 영어 알파벳만으로 구성된 문자열일 경우 들어온 문자열의 길이가 26을 넘을 경우 바로 true를 리턴할 수 있다.
*/

func main() {
	fmt.Println(isDuplicateCharactor("114"))
	fmt.Println(isDuplicateCharactor("145"))
	fmt.Println(isDuplicateCharactor("19151"))
	fmt.Println(isDuplicateCharactor("가나다라마가"))

	fmt.Println("=============================================")

	fmt.Println(isDuplicateCharactor2("114"))
	fmt.Println(isDuplicateCharactor2("145"))
	fmt.Println(isDuplicateCharactor2("19151"))
	fmt.Println(isDuplicateCharactor2("가나다라마가"))

	a := "않"
	b := "안"

	fmt.Println(len(a))
	fmt.Println(len(b))
}

func isDuplicateCharactor(input string) bool {
	var (
		charMap = make(map[string]bool)
		result  = false
	)

	if input == "" {
		return result
	}

	for _, v := range input {
		char := string(v)
		fmt.Println(char, v)

		if _, exist := charMap[char]; exist {
			result = true
			break
		}

		charMap[char] = true
	}

	fmt.Println(charMap)
	return result
}

func isDuplicateCharactor2(input string) bool {
	var result = false

	if input == "" {
		return result
	}

	// len()은 바이트 길이를 반환하기 때문에 한국어와 같은 3byte 문자를 사용하면 깨지게 된다.
	// 이땐 rune타입을 사용하여 문자 하나를 4바이트로 변환시킨 후 출력할 수 있다.

	// for i := 0; i < len(input); i++ {
	// 	fmt.Println(string(input[i]))
	// }

	inputRune := []rune(input)
	for i := 0; i < len(inputRune); i++ {
		char := string(inputRune[i])

		for j := i + 1; j < len(inputRune); j++ {
			char2 := string(inputRune[j])

			if char == char2 {
				return true
			}
		}
	}

	return result
}
