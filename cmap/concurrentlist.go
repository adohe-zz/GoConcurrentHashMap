package cmap

import (
	"unsafe"
	"sync/atomic"
)

var deleteMark = "delete"
var head = "head"

type ListNode struct {
	unsafe.Pointer
	value Any
}

type ConcurrentList struct {
	*ListNode
	size int64
}

func NewConcurrentList() *ConcurrentList {
	return &ConcurrentList{&ListNode{nil, &head}, 0}	
}

func (list *ConcurrentList) Insert(value Any) bool {
		
}

func (list *ConcurrentList) Size() int {
	return int(atomic.LoadInt64(&list.size))
}

func (node *ListNode) next() *ListNode {
	next := atomic.LoadPointer(&node.Pointer)
	
	for next != nil {
		nextNode := (*ListNode)(next)
		
		if sp, ok := nextNode.value.(*string); ok && sp == &deleteMark {
			return nextNode.next()
		}
		
		if nextNode.isDeleted() {
			atomic.CompareAndSwapPointer(&node.Pointer, next, unsafe.Pointer(nextNode.next()))
			next = atomic.LoadPointer(&node.Pointer)
		} else {
			return nextNode
		}
	}
	
	return nil
}

func (node *ListNode) isDeleted() bool {
	next := atomic.LoadPointer(&node.Pointer)
	if next == nil {
		return false
	}
	
	if sp, ok := (*ListNode)(next).value.(*string); ok && sp == &deleteMark {
		return true
	}
	
	return false
}