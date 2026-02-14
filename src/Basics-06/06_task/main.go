package main

import (
	cachePackage "06task/cache"
	semaphore "06task/semaphore"
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	tS := time.Now()
	var tE time.Duration
	fmt.Printf("Start: %v\n", tS)

L:
	for {
		select {
		case <-ctx.Done():
			tE = time.Since(tS)
			fmt.Println("ctx timeout is reached")
			cancel()
			break L
		default:
			fmt.Println("waiting")
			go increaseCacheKey(cache, s)
			go setCacheKey(cache, 3, s)
			time.Sleep(time.Millisecond * 10)
		}
	}

	fmt.Printf("End: %v\n", time.Now())
	fmt.Printf("Done in : %v ms\n", tE)
	fmt.Println(cache.Get(k1))
}

func increaseCacheKey(cache ICache, s *semaphore.Semaphore) {
	s.P()
	fmt.Printf("increaseCacheKey started\n")
	defer s.V()
	defer fmt.Printf("increaseCacheKey finished\n")
	cache.Increase(k1, step)
	time.Sleep(time.Millisecond * 1000)
}

func setCacheKey(cache ICache, i int, s *semaphore.Semaphore) {
	s.P()
	fmt.Printf("setCacheKey started\n")
	defer s.V()
	defer fmt.Printf("setCacheKey finished\n")
	cache.Set(k1, step*i)
	time.Sleep(time.Millisecond * 1000)
}
