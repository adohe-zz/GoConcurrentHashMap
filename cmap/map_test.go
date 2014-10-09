package cmap

import (
	"testing"
)

func TestMapCreation(t *testing.T) {
	m := NewLongConcurrentHashMap()

	if m == nil {
		t.Error("map is null")
	}

	if m.Size() != 0 {
		t.Error("map is not empty")
	}
}
