package main

import (
	"fmt"
)

func main() {
	fmt.Println("beepboop")
	sshift := 0
	ssize := 1
	for ssize < 17 {
		sshift ++
		ssize <<= 1
	}

	c := 33 / ssize
	if c * ssize < 33 {
		c ++
	}
	cap := 1
	for cap < c {
		cap <<= 1
	}
	fmt.Println(ssize)
	fmt.Println(sshift)
	fmt.Println(cap)
}
