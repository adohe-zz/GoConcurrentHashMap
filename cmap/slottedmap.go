package cmap

import (
	"sync"
)

type SlottedLongConcurrentHashMap struct {
	sliceShift int
	sliceMask int
	slices []map[int64]interface{}
	locks []sync.RWMutex
}

func longHash(key int64) int {
	h := int(key ^ (key >> 32))
	h ^= (h >> 20) ^ (h >> 12)
	return h ^ (h >> 7) ^ (h >> 4)
}

func sliceForHash(hash, mask int, shift uint) int {
	return (hash >> shift) & mask
}

func NewSlottedLongConcurrentHashMap(concurrencyLevel, initialCapacity int) *SlottedLongConcurrentHashMap {
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

	for i := range m.slices {
		m.slices[i] = make(map[int64]interface{}, cap)
	}

	return m
}

func (m *SlottedLongConcurrentHashMap) Put(key int64, value interface{}) {
	hash := longHash(key)
	segment := sliceForHash(hash, m.sliceMask, uint(m.sliceShift))

	m.locks[segment].Lock()
	defer m.locks[segment].Unlock()

	m.slices[segment][key] = value
}

func (m *SlottedLongConcurrentHashMap) Get(key int64) (interface{}, bool) {
	hash := longHash(key)
	segment := sliceForHash(hash, m.sliceMask, uint(m.sliceShift))

	m.locks[segment].RLock()
	defer m.locks[segment].RUnlock()

	value, ok := m.slices[segment][key]

	return value, ok
}

func (m *SlottedLongConcurrentHashMap) Clear() {
	for i := range m.slices {
		m.locks[i].Lock()
		m.slices[i] = make(map[int64]interface{})
		m.locks[i].Unlock()
	}	
}

func (m *SlottedLongConcurrentHashMap) Remove(key int64) {
	hash := longHash(key)
	segment := sliceForHash(hash, m.sliceMask, uint(m.sliceShift))
	
	m.locks[segment].Lock()
	defer m.locks[segment].Unlock()
	
	delete(m.slices[segment], key)
}

func (m *SlottedLongConcurrentHashMap) Size() int {
	size := 0
	for _, s := range m.SizeFactors() {
		size += s
	}
	
	return size
}

func (m *SlottedLongConcurrentHashMap) SizeFactors() []int {
	factors := make([]int, m.sliceMask + 1)
	
	for i := range m.slices {
		m.locks[i].RLock()
		factors[i] = len(m.slices[i])
		m.locks[i].RUnlock()
	}
	
	return factors
}

func (m *SlottedLongConcurrentHashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *SlottedLongConcurrentHashMap) Contains(key int64) bool {
	hash := longHash(key)
	segment := sliceForHash(hash, m.sliceMask, uint(m.sliceShift))
	
	m.locks[segment].RLock()
	defer m.locks[segment].RUnlock()
	
	_, ok := m.slices[segment][key]
	
	return ok
}
