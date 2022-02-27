package main

import (
	"fmt"
	"sample/learngo/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	value, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	err = dictionary.Add("first", "111")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success to add")
	}

	err = dictionary.Add("second", "Second word")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success to add")
	}

	err = dictionary.Update("first", "First word")
	if err != nil {
		fmt.Println(err)
	}

	dictionary.Delete("second")

	fmt.Println(dictionary)
}
