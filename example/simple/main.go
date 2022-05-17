package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/mehditeymorian/superchan"
)

const (
	channelSize = 1000

	senderSize           = 10
	senderRandomInterval = false // if true, sender will send messages with random interval in range of [0, 1000] milliseconds
	senderInterval       = 100   // milliseconds

	receiverSize           = 4
	receiverRandomInterval = false // if true, receiver will receive messages with random interval in range of [0, 1000] milliseconds
	receiverInterval       = 100   // millisecond

)

func main() {
	c := superchan.New[int](channelSize)

	for i := 0; i < senderSize; i++ {
		go send(c)
	}

	for i := 0; i < receiverSize; i++ {
		go receive(c)
	}

	for {
		time.Sleep(950 * time.Millisecond)
		log.Printf("buffer: %d in: %d/s out: %d/s\n", c.BufferedSize(), c.InputRate(), c.OutputRate())
	}

}

func send(c *superchan.Chan[int]) {
	for {
		n := senderInterval
		if senderRandomInterval {
			n = rand.Intn(1000)
		}
		c.Send(n)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func receive(c *superchan.Chan[int]) {
	for {
		n := receiverInterval
		if receiverRandomInterval {
			n = rand.Intn(1000)
		}
		_ = c.Receive()
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}
