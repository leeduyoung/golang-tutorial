package mutex

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("[Mutex] create single instance")
			singleInstance = &single{}
		} else {
			fmt.Println("[Mutex] single instance already created-1")
		}
	} else {
		fmt.Println("[Mutex] single instance already created-2")
	}

	return singleInstance
}
