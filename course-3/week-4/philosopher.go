package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Chopstick struct {
	sync.Mutex
}

// Host is a limited reimplementation of golang.org/x/sync/semaphore
// Which implements semaphores in a more sophisticated way
type Host struct {
	cur  int
	size int
	mu   sync.Mutex
}

func NewHost() *Host {
	return &Host{
		cur:  0,
		size: 2, // Max number of philosophers eating
		mu:   sync.Mutex{},
	}
}

func (h *Host) tryAcquire() bool {
	h.mu.Lock()
	success := h.size-h.cur >= 1
	if success {
		h.cur++
	}
	h.mu.Unlock()
	return success
}

func (h *Host) AskForPermission() {
	for !h.tryAcquire() {
		continue
	}
}

func (h *Host) NotifyFinishedEating() {
	h.mu.Lock()
	h.cur--
	h.mu.Unlock()
}

type Philosopher struct {
	host           *Host
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) Eat() {
	for i := 0; i < 3; i++ {
		p.host.AskForPermission()
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("starting to eat %v\n", p.number)
		// Eating for a random duration
		time.Sleep(time.Duration(rand.Intn(10)*100) * time.Millisecond)
		fmt.Printf("finishing eating %v\n", p.number)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		p.host.NotifyFinishedEating()
	}
	wg.Done()
}

func main() {
	host := NewHost()
	chopsticks := make([]*Chopstick, 5)
	for i, _ := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}
	philosophers := make([]*Philosopher, 5)
	for i, _ := range philosophers {
		philosophers[i] = &Philosopher{
			host:           host,
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}
	wg.Add(5)
	for i, _ := range philosophers {
		go philosophers[i].Eat()
	}
	wg.Wait()
}
