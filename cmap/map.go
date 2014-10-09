package cmap

import (
	"sync"
)

type ConcurrentMap struct {
	sync.RWMutex
	m map[int64]interface {}
}
