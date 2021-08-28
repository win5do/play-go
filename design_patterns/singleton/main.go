package main

import (
	"fmt"
	"sync"
)

// Ref: https://www.liwenzhou.com/posts/Go/singleton_in_go/
type singleton struct{}

var s *singleton

// method-1
var lock sync.Mutex

func getInstanceDoubleCheck() *singleton {
	if s == nil {
		lock.Lock()
		defer lock.Unlock()
		if s == nil {
			s = new(singleton)
		}
	}

	return s
}

// method-2
var one sync.Once

func getInstanceSyncOnce() *singleton {
	if s == nil {
		one.Do(func() {
			s = new(singleton)
		})
	}
	return s
}

func main() {
	fmt.Println(getInstanceDoubleCheck())
	fmt.Println(getInstanceSyncOnce())
}
