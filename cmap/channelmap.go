package cmap

type OpCode int

const (
	Get OpCode = iota
	Put
)
type Command struct {
	operation OpCode
	key int64
	value interface {}
	result chan interface {}
}

type ChannelLongConcurrentHashMap struct {
	m map[int64] interface {}
	c chan Command
}

func NewChannelLongConcurrentHashMap(bufferSize int) *ChannelLongConcurrentHashMap {
	m := make(map[int64] interface {})
	c := make(chan Command, bufferSize)

	go func() {

	}()

	return &ChannelLongConcurrentHashMap{m, c}
}
