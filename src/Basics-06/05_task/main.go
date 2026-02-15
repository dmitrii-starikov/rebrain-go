package main

import (
	cachePackage "05task/cache"
	semaphore "05task/semaphore"
	"fmt"
	"time"
)

type ICache interface {
	Set(key string, value int)
	Increase(key string, value int)
	Get(key string) int
	Remove(key string)
}

const (
	k1   = "key1"
	step = 7
)

func main() {
	s := semaphore.NewSemaphore(4)
	cache := cachePackage.NewSafeCache()

	for i := 0; i < 10; i++ {
		s.NewTask()
		go increaseCacheKey(cache, s)
	}

	for i := 0; i < 10; i++ {
		s.NewTask()
		go setCacheKey(cache, i, s)
	}

	for s.IsWorking() {
		time.Sleep(time.Millisecond * 1000)
	}

	fmt.Println(cache.Get(k1))
}

func increaseCacheKey(cache ICache, s *semaphore.Semaphore) {
	s.P()
	fmt.Printf("increaseCacheKey started\n")
	defer s.V()
	cache.Increase(k1, step)
	time.Sleep(time.Millisecond * 1000)
}

func setCacheKey(cache ICache, i int, s *semaphore.Semaphore) {
	s.P()
	fmt.Printf("setCacheKey started\n")
	defer s.V()
	cache.Set(k1, step*i)
	time.Sleep(time.Millisecond * 1000)
}
