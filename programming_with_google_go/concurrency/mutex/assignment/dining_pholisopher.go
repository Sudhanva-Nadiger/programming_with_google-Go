package main

import (
	"fmt"
	"sync"
	"time"
)

var mp map[int]int
var mpMutex sync.Mutex // Add mutex to protect map access

type Host struct {
	// Semaphore to limit concurrent eating to 2 philosophers
	semaphore chan struct{}
}

func NewHost() *Host {
	return &Host{
		semaphore: make(chan struct{}, 2), // Buffer of 2 allows 2 concurrent eaters
	}
}

func (h *Host) GetPermission() {
	h.semaphore <- struct{}{} // Acquire permission
}

func (h *Host) ReleasePermission() {
	<-h.semaphore // Release permission
}

type Stick struct {
	// Mutex to protect the stick
	// and the condition variable
	// to wait for the stick to be available
	// and to signal when the stick is available
	// to the other philosophers
	Mutex *sync.Mutex
}

func (s *Stick) PickUp() {
	// Lock the stick
	s.Mutex.Lock()
}

func (s *Stick) PutDown() {
	// Unlock the stick
	s.Mutex.Unlock()
}

type Philosopher struct {
	Id         int
	LeftStick  *Stick
	RightStick *Stick
	Host       *Host
}

func (p *Philosopher) Eat() {
	mpMutex.Lock()
	if mp[p.Id] >= 3 {
		mpMutex.Unlock()
		return
	}
	mpMutex.Unlock()

	// Get permission from host before eating
	p.Host.GetPermission()

	// Lock the left stick
	p.LeftStick.PickUp()
	// Lock the right stick
	p.RightStick.PickUp()

	fmt.Printf("Starting to eat %v\n", p.Id)

	mpMutex.Lock()
	mp[p.Id]++
	mpMutex.Unlock()

	// Simulate eating time
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Finishing eating %v\n", p.Id)

	// Unlock the right stick
	p.RightStick.PutDown()
	// Unlock the left stick
	p.LeftStick.PutDown()

	// Release permission back to host
	p.Host.ReleasePermission()
}

func main() {
	// Initialize the map
	mp = make(map[int]int)

	// Create host
	host := NewHost()

	// create 5 sticks
	sticks := make([]*Stick, 5)
	for i := range 5 {
		sticks[i] = &Stick{Mutex: &sync.Mutex{}}
	}
	// create 5 philosophers
	philosophers := make([]*Philosopher, 5)
	for i := range 5 {
		philosophers[i] = &Philosopher{
			Id:         i + 1,
			LeftStick:  sticks[i],
			RightStick: sticks[(i+1)%5],
			Host:       host,
		}
	}

	// Wait group to ensure all philosophers finish
	var wg sync.WaitGroup

	// Start all philosophers
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			defer wg.Done()
			for {
				p.Eat()
				mpMutex.Lock()
				if mp[p.Id] >= 3 {
					mpMutex.Unlock()
					return
				}
				mpMutex.Unlock()
			}
		}(philosopher)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating!")
}
