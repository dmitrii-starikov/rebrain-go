package main

import (
	"testing"
)

func BenchmarkWithPool(b *testing.B) {
	counterPool.Put(&Counter{A: 1, B: 2})

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			incrementWithPool()
			b.StartTimer()
		}
	}
}

func BenchmarkWithAllocationButWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// new object every iteration
		c := &Counter{A: i, B: i * 2}

		for j := 0; j < 1000; j++ {
			incrementWithoutPool(c)
		}
	}
}

func BenchmarkWithPool_SingleObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := counterPool.Get().(*Counter)
		c.A = i
		c.B = i * 2

		for j := 0; j < 1000; j++ {
			c.A++
			c.B++
		}

		counterPool.Put(c)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			incrementWithoutPool(&Counter{A: 1, B: 2})
			b.StartTimer()
		}
	}
}
