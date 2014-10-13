package cmap

import (
	"sync"
	"testing"
	"math/rand"
)

func work(n int, m LongHashMap, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		switch OpCode(rand.Intn(3)) {
		case Get:
			m.Get(rand.Int63())
		case Put:
			m.Put(rand.Int63(), rand.Int63())
		case Remove:
			m.Remove(rand.Int63())
		}
	}
	wg.Done()
}

func runBenchmark(n int, m LongHashMap) int {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go work(n/100, m, &wg)
	}
	wg.Wait()
	return m.Size()
}

func BenchmarkSlottedLongMap(b *testing.B) {
	m := NewSlottedLongConcurrentHashMap(16, 16)
	b.Logf("SlottedMap Operations: %d  Length: %d Load:%v\n", b.N, runBenchmark(b.N, m), m.SizeFactors())
}

func BenchmarkSimpleLongMap(b *testing.B) {
	m := NewSimpleLongConcurrentHashMap()
	b.Logf("SlottedMap Operations: %d  Length: %d\n", b.N, runBenchmark(b.N, m))
}
