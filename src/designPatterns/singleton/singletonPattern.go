package singleton

import "sync"

type singleton struct {
}

var s *singleton
var lock sync.Mutex

// method1

func getInstance() *singleton {
	if s == nil {
		lock.Lock()
		defer lock.Unlock()
		if s == nil {
			s = new(singleton)
		}
	}

	return s
}

// method2
var one sync.Once

func getInstance2() *singleton {
	if s == nil {
		one.Do(func() {
			s = new(singleton)
		})
	}

	return s
}
