package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var s *singleton

// method-1
var lock sync.Mutex

func getInstance_lock() *singleton {
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

func getInstance_once() *singleton {
	if s == nil {
		one.Do(func() {
			s = new(singleton)
		})
	}

	return s
}

func main() {
	fmt.Println(getInstance_lock())
	fmt.Println(getInstance_once())
}
