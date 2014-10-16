package cmap

import (
	"fmt"
	"unsafe"
)

const (
	max_exponent = 32
	default_load_factor = 0.75
)

type ConcurrentLongHashMap struct {
	exponent uint32
	buckets []unsafe.Pointer
	size	int64
	loadFactor	float64		
}

func NewConcurrentLongHashMap() *ConcurrentLongHashMap {
	hashMap := &ConcurrentLongHashMap{0, make([]unsafe.Pointer, max_exponent), 0, default_load_factor}
	b := make([]unsafe.Pointer, 1)
	hashMap.buckets[0] = unsafe.Pointer(&b)
	return hashMap
}