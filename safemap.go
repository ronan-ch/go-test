package main

import (
	"strconv"
	"sync"
	"time"
)

type SafeMap struct {
	l      sync.RWMutex
	buffer map[int]string
}

func (p *SafeMap) setValue(key int, value string) {
	p.l.Lock()
	defer p.l.Unlock()
	p.buffer[key] = value
}

func (p *SafeMap) dump() {
	for key, val := range p.buffer {
		println("key= ", key, " val= ", val)
	}
}

func GetMap() *SafeMap {
	return &SafeMap{
		buffer: map[int]string{},
	}
}

func Play() {
	m := GetMap()
	for i := 0; i < 3; i++ {
		go func(v int) {
			m.setValue(v, strconv.Itoa(v))
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	m.dump()
}
