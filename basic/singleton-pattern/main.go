package main

import (
	"singleton/mutex"
	"singleton/once"
	"time"
)

func main() {
	excuteMutexSingleton()
	excuteOnceSingleton()
}

func excuteMutexSingleton() {
	for i := 0; i < 10; i++ {
		go mutex.GetInstance()
	}

	time.Sleep(3 * time.Second)
}

func excuteOnceSingleton() {
	for i := 0; i < 10; i++ {
		go once.GetInstance()
	}

	time.Sleep(3 * time.Second)
}
