package cmap

import (
	"sync"
)

type LongConcurrentHashMap struct {
	sync.RWMutex
	m map[int64]interface {}
}

type MapEntry struct {
	key int64
	value interface {}
}

func NewLongConcurrentHashMap() *LongConcurrentHashMap {
	return &LongConcurrentHashMap{
		m: make(map[int64]interface {}),
	}
}

func (m *LongConcurrentHashMap) Put(key int64, value interface {}) {
	m.Lock()
	defer m.Unlock()

	m.m[key] = value
}

func (m *LongConcurrentHashMap) Get(key int64) (interface {}, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.m[key]

	return value, ok
}

func (m *LongConcurrentHashMap) Remove(key int64) {
	m.Lock()
	defer m.Unlock()

	delete(m.m, key)
}

func (m *LongConcurrentHashMap) Size() int {
	return len(m.m)
}

func (m *LongConcurrentHashMap) Contains(key int64) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.m[key]

	return ok
}

func (m *LongConcurrentHashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *LongConcurrentHashMap) Clear() {
	m.Lock()
	defer m.Unlock()

	m.m = make(map[int64]interface {})
}
