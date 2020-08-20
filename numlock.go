package semaphore

// NumLock is an experimental construct, allowing one to lock a number.
// Primary use case is synchronizing on changing/editing number-based entities such
// as uni-connection TCP ports, or array elements
type NumLock struct {
	mapLock Semaphore
	locks   map[int]*Semaphore
}

// NewNumLock returns an initialized NumLock
func NewNumLock() *NumLock {
	return &NumLock{
		mapLock: NewSemaphore(1),
		locks:   make(map[int]*Semaphore),
	}
}

// Lock locks on the specified number, returning when it is free
func (n *NumLock) Lock(num int) {
	n.mapLock.Lock()
	if _, ok := n.locks[num]; !ok {
		// doesn't exist
		s := NewSemaphore(1)
		n.locks[num] = &s
	}
	n.mapLock.Unlock()
	n.locks[num].Lock()
}

// Unlock unlocks the specified number
func (n *NumLock) Unlock(num int) {
	// no one would call Unlock on a num before calling Lock, right? :(
	n.locks[num].Unlock()
}
