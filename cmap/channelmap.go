package cmap

type OpCode int

const (
	Get OpCode = iota
	Put
	Remove
	Clear
	Size
)
type Command struct {
	operation OpCode
	key int64
	value interface{}
	result chan interface{}
}

type ChannelLongConcurrentHashMap struct {
	m map[int64] interface{}
	c chan Command
}

func NewChannelLongConcurrentHashMap(bufferSize int) *ChannelLongConcurrentHashMap {
	m := make(map[int64] interface{})
	c := make(chan Command, bufferSize)

	go func() {
		for i := range c {
			switch i.operation {
				case Get:
					i.result <- m[i.key]
				case Put:
					m[i.key] = i.value
				case Remove:
					delete(m, i.key)
				case Size:
					
			}
		}
	}()

	return &ChannelLongConcurrentHashMap{m, c}
}

func (m *ChannelLongConcurrentHashMap) Get(key int64) (interface{}, bool) {
	c := Command{Get, key, 0, make(chan interface{})}
	m.c <- c
	
}
