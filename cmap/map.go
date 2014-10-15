package cmap

type OpCode int

const (
	Get OpCode = iota
	Put
	Remove
	Clear
	Size
)

type LongHashMap interface {
	Get(key int64) (interface{}, bool)
	Put(key int64, value interface{})
	Clear()
	Remove(key int64)
	Size() int
	Contains(key int64) bool
	IsEmpty() bool	 
}
