package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := solution("one4seveneight")
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
