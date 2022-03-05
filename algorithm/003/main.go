/*
programmers

숫자 문자열과 영단어
문제 설명
img1.png

네오와 프로도가 숫자놀이를 하고 있습니다. 네오가 프로도에게 숫자를 건넬 때 일부 자릿수를 영단어로 바꾼 카드를 건네주면 프로도는 원래 숫자를 찾는 게임입니다.

다음은 숫자의 일부 자릿수를 영단어로 바꾸는 예시입니다.

1478 → "one4seveneight"
234567 → "23four5six7"
10203 → "1zerotwozero3"
이렇게 숫자의 일부 자릿수가 영단어로 바뀌어졌거나, 혹은 바뀌지 않고 그대로인 문자열 s가 매개변수로 주어집니다. s가 의미하는 원래 숫자를 return 하도록 solution 함수를 완성해주세요.

참고로 각 숫자에 대응되는 영단어는 다음 표와 같습니다.

숫자	영단어
0	zero
1	one
2	two
3	three
4	four
5	five
6	six
7	seven
8	eight
9	nine
제한사항
1 ≤ s의 길이 ≤ 50
s가 "zero" 또는 "0"으로 시작하는 경우는 주어지지 않습니다.
return 값이 1 이상 2,000,000,000 이하의 정수가 되는 올바른 입력만 s로 주어집니다.
입출력 예
s	result
"one4seveneight"	1478
"23four5six7"	234567
"2three45sixseven"	234567
"123"	123
입출력 예 설명
입출력 예 #1

문제 예시와 같습니다.
입출력 예 #2

문제 예시와 같습니다.
입출력 예 #3

"three"는 3, "six"는 6, "seven"은 7에 대응되기 때문에 정답은 입출력 예 #2와 같은 234567이 됩니다.
입출력 예 #2와 #3과 같이 같은 정답을 가리키는 문자열이 여러 가지가 나올 수 있습니다.
입출력 예 #4

s에는 영단어로 바뀐 부분이 없습니다.
제한시간 안내
정확성 테스트 : 10초

*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := solution("one4seveneight")
	fmt.Println(res)

	res = solution2("one4seveneight")
	fmt.Println(res)
}

func solution(s string) int {

	/*
	   0. 영단어가 key, 숫자가 value인 map을 생성한다.
	   1. 문자열을 순회하면서
	   2. 숫자일 경우 pass
	   3. 문자일 경우 문자열 변수에 더한다. 그리고 해당 변수가 map에 존재한다면 숫자로 치환하고 문자열 변수 초기화
	   4. 존재하지 않는다면 다음으로 pass

	   1. dictMap 생성
	   2. 문자열 변수 = strValue 생성
	   3. 리턴할 변수 ret
	*/
	var ret = ""
	var dicMap = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var strValue = ""

	for _, v := range s {
		strV := string(v)
		_, err := strconv.Atoi(strV)

		if err != nil {
			strValue += strV

			dict, exists := dicMap[strValue]
			if exists {
				ret += dict
				strValue = ""
			}

			continue
		}

		ret += strV
	}

	res, _ := strconv.Atoi(ret)
	return res
}

// 간단하게 Replace를 사용하면 된다.
func solution2(s string) int {
	r := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	res, _ := strconv.Atoi(r.Replace(s))
	return res
}
