package cmap

import (
	"sync"
)

type SimpleLongConcurrentHashMap struct {
	sync.RWMutex
	m map[int64]interface {}
}

type MapEntry struct {
	key int64
	value interface {}
}

func NewLongConcurrentHashMap() *SimpleLongConcurrentHashMap {
	return &SimpleLongConcurrentHashMap{
		m: make(map[int64]interface {}),
	}
}

func (m *SimpleLongConcurrentHashMap) Put(key int64, value interface {}) {
	m.Lock()
	defer m.Unlock()

	m.m[key] = value
}

func (m *SimpleLongConcurrentHashMap) Get(key int64) (interface {}, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.m[key]

	return value, ok
}

func (m *SimpleLongConcurrentHashMap) Remove(key int64) {
	m.Lock()
	defer m.Unlock()

	delete(m.m, key)
}

func (m *SimpleLongConcurrentHashMap) Size() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.m)
}

func (m *SimpleLongConcurrentHashMap) Contains(key int64) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.m[key]

	return ok
}

func (m *SimpleLongConcurrentHashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *SimpleLongConcurrentHashMap) Clear() {
	m.Lock()
	defer m.Unlock()

	m.m = make(map[int64]interface {})
}
