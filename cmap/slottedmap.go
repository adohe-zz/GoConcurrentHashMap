package cmap

import (
	"sync"
)

type SlottedLongConcurrentHashMap struct {
	slots int
	segments []map[int64]interface {}
	locks []sync.RWMutex
}

func NewSlottedLongConcurrentHashMap(concurrency uint) *SlottedLongConcurrentHashMap {
	size := 1
	for size < concurrency {
		size <<= 1
	}

	m := &SlottedLongConcurrentHashMap{
		slots: size,
		segments: make(map[[]int64]interface {}, size),
		locks: make([]sync.RWMutex, size),
	}

	for i := range m.segments {
		m.segments[i] = make(map[int64]interface {})
	}

	return m
}

func (m *SlottedLongConcurrentHashMap) Put(key int64, value interface {}) {

}

func (m *SlottedLongConcurrentHashMap) Get(key int64) (interface {}, bool) {

}
