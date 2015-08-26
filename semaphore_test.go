package semaphore

import (
	"testing"
	"time"
)

func TestSemaphore_Lock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(1)

	go func() {
		S.Lock()
		c <- "Lock succeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}

func TestSemaphore_LockUnlock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(1)
	S.Lock()

	go func() {
		S.Unlock()
		c <- "Unlock succeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}

func TestSemaphore_Blocking(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()

	S := NewSemaphore(1)
	S.Lock()

	go func() {
		S.Lock()
		c <- "Lock succeded"
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

func TestSemaphore_BlockingUnlock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(1)
	S.Lock()

	go func() {
		time.Sleep(50 * time.Millisecond)
		S.Unlock()
	}()

	go func() {
		S.Lock()
		c <- "Lock succeded"
	}()

	select {
	case <-c:
		// Good
		return
	case <-timeout:
		t.Error("Lock timed out")
	}
}

func TestSemaphore_BadUnlock(t *testing.T) {

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()

	S := NewSemaphore(1)

	go func() {
		S.Unlock()
		c <- "Unock succeded"
	}()

	select {
	case <-c:
		// Booo
		t.Error("Unlock didn't hold!")
	case <-timeout:
		// Good
		return
	}
}

func TestSemaphore_Free(t *testing.T) {

	S := NewSemaphore(10)

	if S.Free() != 10 {
		t.Errorf("Free should be 10, but is %d!\n", S.Free())
	}

	go func() {
		S.Lock()
		defer S.Unlock()
		time.Sleep(200 * time.Millisecond)
	}()

	time.Sleep(50 * time.Millisecond)
	if S.Free() != 9 {
		t.Errorf("Free should be 9, but is %d!\n", S.Free())
	}

	go func() {
		for i := 0; i < 9; i++ {
			S.Lock()
			defer S.Unlock()
		}
		time.Sleep(200 * time.Millisecond)
	}()

	time.Sleep(50 * time.Millisecond)
	if S.Free() != 0 {
		t.Errorf("Free should be 0, but is %d!\n", S.Free())
	}

	time.Sleep(200 * time.Millisecond)
	if S.Free() != 10 {
		t.Errorf("Ending free should be 10, but is %d!\n", S.Free())
	}

}
