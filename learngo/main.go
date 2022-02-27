package main

import "fmt"

type Person struct {
	name    string
	age     int
	favFood []string
}

func main() {

	names := []string{"nico", "kaye", "mango"}
	names = append(names, "flynn")
	fmt.Println(names)

	tmap := map[string]string{
		"name": "kaye",
		"age":  "one",
	}
	fmt.Println(tmap)

	for key, value := range tmap {
		fmt.Println(key, value)
	}

	nico := Person{
		"nico", 18, []string{"kimchi", "ramen"},
	}

	fmt.Println(nico)
	a := nico.favFood
	fmt.Println(a)
}
