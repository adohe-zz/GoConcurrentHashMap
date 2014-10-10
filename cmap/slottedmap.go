package cmap

import (
	"sync"
)

type SlottedLongConcurrentHashMap struct {
	sliceShift uint
	sliceMask uint
	slices []map[int64]interface {}
	locks []sync.RWMutex
}

func longHash(key int64) int {
	h := (key ^ (key >>> 32)).(int)
	h ^= (h >>> 20) ^ (h >>> 12)
	return h ^ (h >>> 7) ^ (h >>> 4)
}

func NewSlottedLongConcurrentHashMap(concurrencyLevel, initialCapacity uint) *SlottedLongConcurrentHashMap {
	sshift := 0
	ssize := 1
	for ssize < concurrencyLevel {
		sshift ++
		ssize <<= 1
	}

	m := &SlottedLongConcurrentHashMap{
		sliceShift: 32 - sshift,
		sliceMask: ssize - 1,
		slices: make([]map[int64]interface {}, ssize),
		locks: make([]sync.RWMutex, ssize),
	}

	c := initialCapacity / ssize
	if c * ssize < initialCapacity {
		c ++
	}
	cap := 1
	for cap < c {
		cap <<= 1
	}

	for i := range m.segments {
		m.segments[i] = make(map[int64]interface {}, cap)
	}

	return m
}

func (m *SlottedLongConcurrentHashMap) Put(key int64, value interface {}) {
	hash := longHash(key)
	segment := (hash >>> m.sliceShift) & m.sliceMask

	m.locks[segment].Lock()
	defer m.locks[segment].Unlock()

	m.segments[segment][key] = value
}

func (m *SlottedLongConcurrentHashMap) Get(key int64) (interface {}, bool) {
	hash := longHash(key)
	segment := (hash >>> m.sliceShift) & m.sliceMask

	m.locks[segment].RLock()
	defer m.locks[segment].RUnlock()

	value, ok := m.segments[segment][key]

	return value, ok
}
