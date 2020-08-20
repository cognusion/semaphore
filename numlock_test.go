package semaphore

import (
	"testing"
	"time"
)

func TestNumLock_Lock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewNumLock()

	go func() {
		S.Lock(66)
		c <- "Lock succeeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}

func TestNumLock_LockUnlock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewNumLock()
	S.Lock(66)

	go func() {
		S.Unlock(66)
		c <- "Unlock succeeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}

func TestNumLock_Blocking(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()

	S := NewNumLock()
	S.Lock(66)

	go func() {
		S.Lock(66)
		c <- "Lock succeeded"
	}()

	select {
	case r := <-c:
		// Boooo
		t.Errorf("Lock didn't hold! '%s'\n", r)
	case <-timeout:
		// Timed out, as it should
		return
	}
}

func TestNumLock_BlockingUnlock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	S := NewNumLock()
	S.Lock(66)

	go func() {
		time.Sleep(50 * time.Millisecond)
		S.Unlock(66)
	}()

	go func() {
		S.Lock(66)
		c <- "Lock succeeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}
