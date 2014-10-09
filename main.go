package main

import (
	"fmt"
	"cmap"
)

func main() {
	fmt.Println("beepboop")
	m := cmap.NewLongConcurrentHashMap()

	m.Put(5, "hello")

	temp, ok := m.Get(5)

	if ok == true {
		value := temp.(string)
		fmt.Println(value)
	}

	m.Remove(5)

	fmt.Println("helloworld")
}
