// Package semaphore is a super simple goro-safe semaphore struct for Go.
package semaphore

import (
	"fmt"
	"time"
)

// UntilFreeTimeout is a Duration for the goro spawned by the Until func to wait
// if there is nothing consuming messages from provided chan, before releasing
// the lock.
const UntilFreeTimeout = 2 * time.Millisecond

// Semaphore is a goro-safe simple semaphore
type Semaphore struct {
	lock chan bool
}

// NewSemaphore returns a Semaphore allowing up to 'size' locks before blocking
func NewSemaphore(size int) Semaphore {
	return Semaphore{
		lock: make(chan bool, size),
	}
}

// Until returns a channel that fires bool(true) when the lock can be consumed.
// If nothing is listening to the returned channel, after 2ms (UntilFreeTimeout)
// the lock will be removed, so if you're not paying attention, don't use this function.
func (s *Semaphore) Until() <-chan bool {
	b := make(chan bool)
	go func(b chan bool) {
		s.lock <- true
		// We have a lock, let's make sure we should keep it
		select {
		case b <- true:
			// The other side received, assume it's valid.
			return
		case <-time.After(UntilFreeTimeout):
			// The other side did not receive, assume they are gone.
			<-s.lock
			return
		}
	}(b)
	return b
}

// Lock consumes a lock in the semaphore, blocking if none is available
func (s *Semaphore) Lock() {
	s.lock <- true
}

// Unlock replaces a lock in the semaphore, blocking if no locks are consumed
func (s *Semaphore) Unlock() {
	<-s.lock
}

// Add consumes numLocks locks in the semaphore, blocking if none is available
func (s *Semaphore) Add(numLocks int) {
	for i := 0; i < numLocks; i++ {
		s.lock <- true
	}
}

// Sub replaces numLocks locks in the semaphore, blocking if no locks are consumed
func (s *Semaphore) Sub(numLocks int) {
	for i := 0; i < numLocks; i++ {
		<-s.lock
	}
}

// Free returns the number of available locks in the semaphore
func (s *Semaphore) Free() int {
	return cap(s.lock) - len(s.lock)
}

// IsFree takes a Duration, and makes a decent try on determining if someone consumed
// a lock over the Duration, ala a WaitGroup.Wait().
func (s *Semaphore) IsFree(freeFor time.Duration) <-chan bool {
	// Logic is that if we get two consencutive "empty" channels freeFor/2 apart,
	// we consider it done.
	b := make(chan bool, 1)
	go func(b chan bool) {
		halfTime := freeFor / 2
		c := 0
		for {
			<-time.After(halfTime)
			if len(s.lock) == 0 && c > 0 {
				b <- true
				return
			} else if len(s.lock) == 0 {
				c++
			} else if c > 0 {
				c--
			}
		}
	}(b)
	return b
}

// String returns the string representation of the semaphore
func (s *Semaphore) String() string {
	return fmt.Sprintf("%d of %d free", s.Free(), cap(s.lock))
}
