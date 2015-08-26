# semaphore
Very simple goro-safe semaphore

Basics
======

```bash
go get github.com/cognusion/semaphore
```

```go

import (
	"github.com/cognusion/semaphore"
	"time"
	"fmt"
)

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

time.Sleep(3 * time.Second)
```
