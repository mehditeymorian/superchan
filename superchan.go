package superchan

import (
	"sync/atomic"
	"time"
)

const updateDuration = time.Duration(1000) * time.Millisecond

type Chan[T any] struct {
	data    chan T
	inRate  uint32
	outRate uint32
}

func New[T any](size int) *Chan[T] {
	c := &Chan[T]{
		data:    make(chan T, size),
		inRate:  0,
		outRate: 0,
	}
	go c.calculateRate()

	return c
}

func (c *Chan[T]) Send(data T) {
	atomic.AddUint32(&c.inRate, 1)
	c.data <- data
}

func (c *Chan[T]) Receive() T {
	t := <-c.data
	atomic.AddUint32(&c.outRate, 1)

	return t
}

func (c *Chan[T]) InputRate() uint32 {
	return atomic.LoadUint32(&c.inRate)
}

func (c *Chan[T]) OutputRate() uint32 {
	return atomic.LoadUint32(&c.outRate)
}

// BufferedSize returns the number of elements in the channel.
func (c *Chan[T]) BufferedSize() int {
	return len(c.data)
}

func (c *Chan[T]) calculateRate() {
	for range time.Tick(updateDuration) {
		atomic.StoreUint32(&c.inRate, 0)
		atomic.StoreUint32(&c.outRate, 0)
	}
}
