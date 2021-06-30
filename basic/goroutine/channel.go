package main

func sum(a int, b int, c chan int) {
	c <- a + b
}
