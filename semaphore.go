/*
package main

import (
	"github.com/cognusion/semaphore"
	"time"
	"fmt"
)

func main() {
	// Make a new semaphore, with the number of
	// simultaneous locks you want to allow
	S := NewSemaphore(1)
	
	go func() {
		// Call lock, which will block if there aren't free locks
		// and defer the unlock until the function ends
		S.Lock()
		defer S.Unlock()
	
		// Do some stuff
		fmt.Println("Doing some stuff")
		time.Sleep(1 * time.Second)
	}()
	
	go func() {
		// Call lock, which will block if there aren't free locks
		// and defer the unlock until the function ends
		S.Lock()
		defer S.Unlock()
	
		// Do some other stuff
		fmt.Println("Doing some other stuff")
		time.Sleep(50 * time.Millisecond)
	}()
	
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Free locks? %d\n",S.Free())
	time.Sleep(3 * time.Second)
	fmt.Printf("Free locks now? %d\n",S.Free())
}
*/
package semaphore

type Semaphore struct {
	lock chan bool
}

// Returns a Semaphore allowing up to 'size' locks before blocking
func NewSemaphore(size int) Semaphore {
	var S Semaphore
	S.lock = make(chan bool, size)
	return S
}

// Consume a lock in the semaphore, blocking if none is available
func (s *Semaphore) Lock() {
	s.lock <- true
}

// Replace a lock in the semaphore, blocking if no locks are consumed
func (s *Semaphore) Unlock() {
	<-s.lock
}

// Return the number of available locks in the semaphore
func (s *Semaphore) Free() int {
	return cap(s.lock) - len(s.lock)
}
