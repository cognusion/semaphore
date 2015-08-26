package semaphore

type Semaphore struct {
	lock chan bool
}

func NewSemaphore(size int) Semaphore {
	var S Semaphore
	S.lock = make(chan bool, size)
	return S
}

func (s *Semaphore) Lock() {
	s.lock <- true
}

func (s *Semaphore) Unlock() {
	<-s.lock
}
