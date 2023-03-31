package main

import (
	"sync"
)

type ChanKey interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type Channel[T ChanKey] struct {
	mut    sync.Mutex
	C      chan T
	closed bool
}

func NewChannel[T ChanKey](size int) *Channel[T] {
	return &Channel[T]{
		C:      make(chan T, size),
		closed: false,
		mut:    sync.Mutex{},
	}
}

func (c *Channel[T]) Close() {
	c.mut.Lock()
	defer c.mut.Unlock()
	if !c.closed {
		close(c.C)
		c.closed = true
	}
}

func (c *Channel[T]) IsClosed() bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.closed
}

func main2() {
	ch := NewChannel[int](2)
	ch.C <- 1
	println(ch.IsClosed())
	ch.Close()
	ch.Close()
	println(ch.IsClosed())
}
