package cmap

import (
	"fmt"
	"unsafe"
)

const (
	loadFactor = 0.75
)

type ConcurrentLongHashMap struct {
	buckets []unsafe.Pointer
	size	int64
	loadFactor	float64		
}

func NewConcurrentLongHashMap() *ConcurrentLongHashMap {
	hashMap := &NewConcurrentLongHashMap{}
	
	return hashMap
}