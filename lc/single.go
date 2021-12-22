package lc

import "sync"

type single struct {
}

var s *single
var mu *sync.Mutex
var once sync.Once

//懒汉模式
func getInstance() *single {
	if s == nil {
		s = &single{}
	}
	return s
}

//带锁的单例模式
func getInstance1() *single {
	mu.Lock()
	defer mu.Unlock()

	if s == nil {
		s = &single{}
	}

	return s
}

//双重锁
func getInstance2() *single {
	if s == nil {
		mu.Lock()
		defer mu.Unlock()
		if s == nil {
			s = &single{}
		}
	}

	return s
}

//sync.Once
func getInstance3() *single {
	once.Do(func() {
		s = &single{}
	})

	return s
}
