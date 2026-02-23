package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	A int
	B int
}

var counterPool = sync.Pool{
	New: func() any { return &Counter{} },
}

func incrementWithPool() {
	c := counterPool.Get()
	if c != nil {
		counter := c.(*Counter)
		counter.A++
		counter.B++
		counterPool.Put(counter)
	}
}

func incrementWithoutPool(counter *Counter) {
	counter.A++
	counter.B++
}

func main() {
	counterPool.Put(&Counter{A: 1, B: 2})

	for i := 0; i < 100; i++ {
		incrementWithPool()
	}

	result := counterPool.Get().(*Counter)
	fmt.Printf("Result 1: Counter{A: %d, B: %d}\n", result.A, result.B)

	//////////////////////
	counterPool.Put(&Counter{A: 100, B: 200})

	for i := 0; i < 50; i++ {
		incrementWithPool()
	}

	result = counterPool.Get().(*Counter)
	fmt.Printf("Result 2: Counter{A: %d, B: %d}\n", result.A, result.B)

	//////////////////////
	c := counterPool.Get().(*Counter)
	c.A, c.B = 123, 321

	for j := 0; j < 1000; j++ {
		c.A++
		c.B++
	}

	counterPool.Put(c)

	result = counterPool.Get().(*Counter)
	fmt.Printf("Result 3: Counter{A: %d, B: %d}\n", result.A, result.B)

	//////////////////////
	ctr := &Counter{A: 1000, B: 8888}

	for j := 0; j < 1000; j++ {
		incrementWithoutPool(ctr)
	}
	fmt.Printf("Result 4: Counter{A: %d, B: %d}\n", ctr.A, ctr.B)
}
