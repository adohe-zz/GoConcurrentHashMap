package cmap

type OpCode int

type Any interface{}

type Comparable interface {
	CompareTo(Any) int	
}

type Equable interface {
	Equals(Any) bool
}

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
