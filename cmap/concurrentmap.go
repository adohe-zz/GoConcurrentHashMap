package cmap

import (
	"fmt"
	"unsafe"
)

const (
	max_exponent = 32
	default_load_factor = 0.75
)

type WrapperKey interface {
	Equable
	HashCode() int
}

type Node struct {
	hashCode int
	hashKey int
	key WrapperKey
	value unsafe.Pointer
}

func newRealNodeWithHashCode(key WrapperKey, value Any, hashCode int) *Node {
	return &Node{hashCode, , key, unsafe.Pointer(&value)}
}

func newRealNode(key WrapperKey, value Any) *Node {
	return newRealNodeWithHashCode(key, value, key.HashCode())
}

type LongKey int64

func (key LongKey) HashCode() int {
	h := int(key ^ (key >> 32))
	h ^= (h >> 20) ^ (h >> 12)
	return h ^ (h >> 7) ^ (h >> 4)
}

func (key LongKey) Equals(any Any) bool {
	if lk, ok := any.(LongKey); ok {
		return int64(key) == int64(lk)
	}	
	
	return false
}

type ConcurrentLongHashMap struct {
	exponent uint32
	buckets []unsafe.Pointer
	size	 int64
	loadFactor	float64		
}

func NewConcurrentLongHashMap() *ConcurrentLongHashMap {
	hashMap := &ConcurrentLongHashMap{0, make([]unsafe.Pointer, max_exponent), 0, default_load_factor}
	b := make([]unsafe.Pointer, 1)
	hashMap.buckets[0] = unsafe.Pointer(&b)
	return hashMap
}

func (m *ConcurrentLongHashMap) GetByHashCode(hashCode int, key WrapperKey) (any Any, ok bool) {
	
}

func (m *ConcurrentLongHashMap) Get(key WrapperKey) (Any, bool) {
	
}