package once

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("[Once] create single instance")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("[Once] single instance already created-1")
	}
	return singleInstance
}
