package cmap

import (
	"unsafe"
)

type ListNode struct {
	unsafe.Pointer
	value Any
}