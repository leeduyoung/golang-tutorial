package main

import (
	"fmt"
	"math"
	"strings"
)

/*
1 Failed) 조이스틱 (Greedy - 탐욕법)
문제 설명
조이스틱으로 알파벳 이름을 완성하세요. 맨 처음엔 A로만 이루어져 있습니다.
ex) 완성해야 하는 이름이 세 글자면 AAA, 네 글자면 AAAA

조이스틱을 각 방향으로 움직이면 아래와 같습니다.

▲ - 다음 알파벳
▼ - 이전 알파벳 (A에서 아래쪽으로 이동하면 Z로)
◀ - 커서를 왼쪽으로 이동 (첫 번째 위치에서 왼쪽으로 이동하면 마지막 문자에 커서)
▶ - 커서를 오른쪽으로 이동 (마지막 위치에서 오른쪽으로 이동하면 첫 번째 문자에 커서)
예를 들어 아래의 방법으로 "JAZ"를 만들 수 있습니다.

- 첫 번째 위치에서 조이스틱을 위로 9번 조작하여 J를 완성합니다.
- 조이스틱을 왼쪽으로 1번 조작하여 커서를 마지막 문자 위치로 이동시킵니다.
- 마지막 위치에서 조이스틱을 아래로 1번 조작하여 Z를 완성합니다.
따라서 11번 이동시켜 "JAZ"를 만들 수 있고, 이때가 최소 이동입니다.
만들고자 하는 이름 name이 매개변수로 주어질 때, 이름에 대해 조이스틱 조작 횟수의 최솟값을 return 하도록 solution 함수를 만드세요.

제한 사항
name은 알파벳 대문자로만 이루어져 있습니다.
name의 길이는 1 이상 20 이하입니다.
입출력 예
name	return
"JEROEN"	56
"JAN"	23
출처

※ 공지 - 2019년 2월 28일 테스트케이스가 추가되었습니다.
※ 공지 - 2022년 1월 14일 지문 수정 및 테스트케이스가 추가되었습니다. 이로 인해 이전에 통과하던 코드가 더 이상 통과하지 않을 수 있습니다.
*/

var alpabat = []string{
	"A", "B", "C", "D", "E",
	"F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y",
	"Z",
}

func main() {
	fmt.Println("solution: ", solution("JEROEN")) // 56
	fmt.Println("solution: ", solution("JAN"))    // 23
	fmt.Println("solution: ", solution("JAJAAJ")) // 32

	// ============= //

	fmt.Println("solution2: ", solution2("JEROEN")) // 56
	fmt.Println("solution2: ", solution2("JAN"))    // 23
	fmt.Println("solution2: ", solution2("JAJAAJ")) // 31

	// fmt.Println(solution("JEROEN") == 56)
	// fmt.Println(solution("JAN") == 23)
}

func solution(name string) int {
	sum := 0
	arr := []string{}
	aCount := 0

	// TODO: How to find distance?
	// 가장 가까운 A가 아닌 문자를 찾는다.

	for _, v := range name {
		c := fmt.Sprintf("%c", v)

		if c == "A" {
			// aCount++
			continue
		}

		arr = append(arr, c)
	}

	if aCount == 0 {
		sum += len(name) - 1
	}

	for _, v := range arr {
		upCount := getUpCount(v)
		downCount := getDownCount(v)

		if upCount < downCount {
			sum += upCount
			continue
		}

		sum += downCount
	}

	return sum
}

func getUpCount(c string) int {
	for i, v := range alpabat {
		if v == c {
			return i
		}
	}
	return 0
}

func getDownCount(c string) int {
	for i := range alpabat {
		if c == alpabat[len(alpabat)-1-i] {
			return i + 1
		}
	}
	return 0
}

// solution2
func solution2(name string) int {
	var answer = 0
	var length = len(name)
	var upDownCount = 0
	var moveCount = length - 1 // 한 방향으로 쭉 갔을때

	// UP & DOWN COUNT
	for i := 0; i < length; i++ {
		upDownCount += minUpOrDownCount(fmt.Sprintf("%c", name[i]))
	}

	// LEFT & RIGHT COUNT
	for startOfA := 0; startOfA < length; startOfA++ {
		endOfA := startOfA + 1
		for (endOfA < length) && (fmt.Sprintf("%c", name[endOfA]) == "A") {
			endOfA++
		}

		moveCount = int(math.Min(float64(moveCount), float64(startOfA*2+length-endOfA)))
		moveCount = int(math.Min(float64(moveCount), float64((length-endOfA)*2+startOfA)))
	}

	answer = upDownCount + moveCount
	return answer
}

func minUpOrDownCount(c string) int {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index := strings.Index(alphabet, c)
	if index == -1 {
		fmt.Printf("not found alphabet, input: %s", c)
	}
	return int(math.Min(float64(index), float64(len(alphabet)-index)))
}
