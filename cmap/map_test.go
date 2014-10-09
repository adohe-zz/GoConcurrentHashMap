package cmap

import (
	"testing"
	"strconv"
)

type Girl struct {
	name string
}

func TestMapCreation(t *testing.T) {
	m := NewLongConcurrentHashMap()

	if m == nil {
		t.Error("map is null")
	}

	if m.Size() != 0 {
		t.Error("map is not empty")
	}
}

func TestPut(t *testing.T) {
	m := NewLongConcurrentHashMap()

   	lily := Girl{"Lily"}
	lucy := Girl{"Lucy"}

	m.Put(2, lily)
	m.Put(3, lucy)

	if m.Size() != 2 {
		t.Error("map should contain only two elements")
	}
}

func TestGet(t *testing.T) {
	m := NewLongConcurrentHashMap()

	value, ok := m.Get(1)

	if ok == true {
		t.Error("ok should be false")
	}
	if value != nil {
		t.Error("value should be nil")
	}

	clair := Girl{"Clair"}

	m.Put(1, clair)
	temp, ok := m.Get(1)

	value := temp.(Girl)

	if ok == false {
		t.Error("ok should be true")
	}
	if &value == nil {
		t.Error("value should not be null")
	}
	if value.name != "Clair" {
		t.Error("value is modified")
	}
}

func TestRemove(t *testing.T) {
	m := NewLongConcurrentHashMap()

	alice := Girl{"Alice"}
	m.Put(3, alice)

	temp, ok := m.Get(3)
	if ok == false {
		t.Error("ok should be true")
	}

	m.Remove(3)
	temp, ok := m.Get(3)
	if ok == true {
		t.Error("ok should be false")
	}
	if temp != nil {
		t.Error("temp should be null")
	}
}

func TestSize(t *testing.T) {
	m := NewLongConcurrentHashMap()

	if m.Size() != 0 {
		t.Error("map should be empty")
	}

	alma := Girl{"Alma"}
	alva := Girl{"Alva"}
	m.Put(1, alma)
	m.Put(2, alva)

	if m.Size() != 2 {
		t.Error("map should just contain only two elements")
	}
}

func TestContains(t *testing.T) {
	m := NewLongConcurrentHashMap()

	if m.Contains(1) == true {
		t.Error("map should not contain this key")
	}

	amy := Girl{"Amy"}
	m.Put(7, amy)

	if m.Contains(7) == false {
		t.Error("map should contain this key")
	}
}

func TestIsEmpty(t *testing.T) {
	m := NewLongConcurrentHashMap()

	if m.IsEmpty() == false {
		t.Error("map should be empty")
	}

	m.Put(1, Girl{"Andrea"})
	if m.IsEmpty() == true {
		t.Error("map should not be empty")
	}
}

func TestClear(t *testing.T) {
	m := NewLongConcurrentHashMap()

	m.Clear()
	if m.Size() != 0 {
		t.Error("map should be empty")
	}

	m.Put(1, Girl{"Amanda"})
	m.Clear()
	if m.Size() != 0 {
		t.Error("expect an emtpy map")
	}
}

func TestConcurrent(t *testing.T) {
	m := NewLongConcurrentHashMap()
	ch := make(chan string)
	const loop = 2000
	var s [loop]string

	go func() {
		for i := 0; i < loop/2; i++ {
			m.Put(i, strconv.Itoa(i))

			value, _ := m.Get(i)

			ch <- value.(string)
		}
	}()

	go func() {
		for i := loop/2; i < loop; i++ {
			m.Put(i, strconv.Itoa(i))

			value, _ := m.Get(i)

			ch <- value.(string)
		}
	}()

	// wait
	counter := 0
	for ele := range ch {
		s[counter] = ele;
		counter ++;
		if counter == loop {
			break;
		}
	}

	if m.Size() != loop {
		t.Error("map should contain 2000 elements")
	}
}
