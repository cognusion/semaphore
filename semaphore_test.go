package semaphore

import (
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
)

func TestSemaphore_Lock(t *testing.T) {
	defer leaktest.Check(t)()

	timeout := make(chan bool, 1)
	c := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(1)

	go func() {
		S.Lock()
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

func TestSemaphore_Add(t *testing.T) {
	defer leaktest.Check(t)()

	timeout := make(chan bool, 1)
	c := make(chan string, 4)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(4)

	go func() {
		before := S.Free()
		if before != 4 {
			t.Errorf("Before Add() Free should have been 4, but was %d\n", before)
		}

		S.Add(4)

		after := S.Free()
		if after != 0 {
			t.Errorf("After Add() Free should have been 0, but was %d\n", after)
		}

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

func TestSemaphore_AddSub(t *testing.T) {
	defer leaktest.Check(t)()

	timeout := make(chan bool, 1)
	c := make(chan string, 4)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	S := NewSemaphore(4)

	go func() {
		before := S.Free()
		if before != 4 {
			t.Errorf("Before Add() Free should have been 4, but was %d\n", before)
		}

		S.Add(4)

		after := S.Free()
		if after != 0 {
			t.Errorf("After Add() Free should have been 0, but was %d\n", after)
		}

		c <- "Lock succeeded"
	}()

	select {
	case <-c:
		// Good
	case <-timeout:
		t.Error("Lock timed out")
	}

	go func() {
		before := S.Free()
		if before != 0 {
			t.Errorf("Before Sub() Free should have been 0, but was %d\n", before)
		}

		S.Sub(4)

		after := S.Free()
		if after != 4 {
			t.Errorf("After Sub() Free should have been 4, but was %d\n", after)
		}

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

func TestSemaphore_LockUnlock(t *testing.T) {
	defer leaktest.Check(t)()

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

func TestSemaphore_UntilLockUnlock(t *testing.T) {
	defer leaktest.Check(t)()

	S := NewSemaphore(1)
	S.Lock()

	go func() {
		<-time.After(100 * time.Millisecond)
		S.Unlock()
	}()

	select {
	case <-S.Until():
		// Good
		return
	case <-time.After(1 * time.Second):
		// Timeout, bad!
		t.Error("Lock timed out")
	}
}

func TestSemaphore_Blocking(t *testing.T) {
	defer leaktest.Check(t)()

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
		c <- "Lock succeeded"
	}()

	select {
	case r := <-c:
		// Boooo
		t.Errorf("Lock didn't hold! '%s'\n", r)
	case <-timeout:
		// Timed out, as it should
		defer S.Unlock()
		return
	}
}

func TestSemaphore_BlockingUnlock(t *testing.T) {
	defer leaktest.Check(t)()

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

func TestSemaphore_BadUnlock(t *testing.T) {
	defer leaktest.Check(t)()

	c := make(chan string, 1)

	S := NewSemaphore(1)

	go func() {
		S.Unlock()
		c <- "Unlock succeeded"
	}()

	select {
	case <-c:
		// Booo
		t.Error("Unlock didn't hold!")
	case <-time.After(100 * time.Millisecond):
		// Good
		defer S.Lock() // haha
		return
	}
}

func TestSemaphore_Free(t *testing.T) {
	defer leaktest.Check(t)()

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

func BenchmarkSemaphore1k(b *testing.B) {
	ssize := 1000

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		S := NewSemaphore(ssize)
		go func() {
			for i := 0; i < ssize; i++ {
				S.Lock()
				defer S.Unlock()
			}
		}()
		for {
			// WTF?
			if S.Free() == ssize {
				break
			}
		}
	}
}
