

# semaphore
`import "github.com/cognusion/semaphore"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package semaphore is a super simple goro-safe semaphore struct for Go.
* NewSemaphore(N) to create a semaphore of size N
* Lock() to consume
* Unlock() to replace
* Add(i) to add i to the lock count
* Sub(i) to subtract i to the lock count
* Free() to see how many locks are available
* <-Until() to wait until a channel get a message to consume.


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




## <a name="pkg-index">Index</a>
* [type Counter](#Counter)
  * [func NewCounter() Counter](#NewCounter)
  * [func NewSetCounter(value int) Counter](#NewSetCounter)
  * [func (c *Counter) Dec()](#Counter.Dec)
  * [func (c *Counter) Inc()](#Counter.Inc)
  * [func (c *Counter) Value() int](#Counter.Value)
* [type NumLock](#NumLock)
  * [func NewNumLock() *NumLock](#NewNumLock)
  * [func (n *NumLock) Lock(num int)](#NumLock.Lock)
  * [func (n *NumLock) Unlock(num int)](#NumLock.Unlock)
* [type Semaphore](#Semaphore)
  * [func NewSemaphore(size int) Semaphore](#NewSemaphore)
  * [func (s *Semaphore) Add(numLocks int)](#Semaphore.Add)
  * [func (s *Semaphore) Free() int](#Semaphore.Free)
  * [func (s *Semaphore) Lock()](#Semaphore.Lock)
  * [func (s *Semaphore) String() string](#Semaphore.String)
  * [func (s *Semaphore) Sub(numLocks int)](#Semaphore.Sub)
  * [func (s *Semaphore) Unlock()](#Semaphore.Unlock)
  * [func (s *Semaphore) Until() &lt;-chan bool](#Semaphore.Until)


#### <a name="pkg-files">Package files</a>
[counter.go](https://github.com/cognusion/semaphore/tree/master/counter.go) [numlock.go](https://github.com/cognusion/semaphore/tree/master/numlock.go) [semaphore.go](https://github.com/cognusion/semaphore/tree/master/semaphore.go)






## <a name="Counter">type</a> [Counter](https://github.com/cognusion/semaphore/tree/master/counter.go?s=60:107#L4)
``` go
type Counter struct {
    Semaphore
    // contains filtered or unexported fields
}

```
Counter is a Semaphore-locked integer







### <a name="NewCounter">func</a> [NewCounter](https://github.com/cognusion/semaphore/tree/master/counter.go?s=443:468#L29)
``` go
func NewCounter() Counter
```
NewCounter returns a new Counter, with a 0 value


### <a name="NewSetCounter">func</a> [NewSetCounter](https://github.com/cognusion/semaphore/tree/master/counter.go?s=563:600#L34)
``` go
func NewSetCounter(value int) Counter
```
NewSetCounter returns a new Counter with the specified value





### <a name="Counter.Dec">func</a> (\*Counter) [Dec](https://github.com/cognusion/semaphore/tree/master/counter.go?s=231:254#L17)
``` go
func (c *Counter) Dec()
```
Dec [rements] the counter




### <a name="Counter.Inc">func</a> (\*Counter) [Inc](https://github.com/cognusion/semaphore/tree/master/counter.go?s=138:161#L10)
``` go
func (c *Counter) Inc()
```
Inc [rements] the counter




### <a name="Counter.Value">func</a> (\*Counter) [Value](https://github.com/cognusion/semaphore/tree/master/counter.go?s=338:367#L24)
``` go
func (c *Counter) Value() int
```
Value returns the current counter value




## <a name="NumLock">type</a> [NumLock](https://github.com/cognusion/semaphore/tree/master/numlock.go?s=225:295#L6)
``` go
type NumLock struct {
    // contains filtered or unexported fields
}

```
NumLock is an experimental construct, allowing one to lock a number.
Primary use case is synchronizing on changing/editing number-based entities such
as uni-connection TCP ports, or array elements







### <a name="NewNumLock">func</a> [NewNumLock](https://github.com/cognusion/semaphore/tree/master/numlock.go?s=342:368#L12)
``` go
func NewNumLock() *NumLock
```
NewNumLock returns an initialized NumLock





### <a name="NumLock.Lock">func</a> (\*NumLock) [Lock](https://github.com/cognusion/semaphore/tree/master/numlock.go?s=525:556#L20)
``` go
func (n *NumLock) Lock(num int)
```
Lock locks on the specified number, returning when it is free




### <a name="NumLock.Unlock">func</a> (\*NumLock) [Unlock](https://github.com/cognusion/semaphore/tree/master/numlock.go?s=758:791#L32)
``` go
func (n *NumLock) Unlock(num int)
```
Unlock unlocks the specified number




## <a name="Semaphore">type</a> [Semaphore](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1303:1344#L55)
``` go
type Semaphore struct {
    // contains filtered or unexported fields
}

```
Semaphore is a goro-safe simple semaphore







### <a name="NewSemaphore">func</a> [NewSemaphore](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1426:1463#L60)
``` go
func NewSemaphore(size int) Semaphore
```
NewSemaphore returns a Semaphore allowing up to 'size' locks before blocking





### <a name="Semaphore.Add">func</a> (\*Semaphore) [Add](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2062:2099#L87)
``` go
func (s *Semaphore) Add(numLocks int)
```
Add consumes numLocks locks in the semaphore, blocking if none is available




### <a name="Semaphore.Free">func</a> (\*Semaphore) [Free](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2394:2424#L101)
``` go
func (s *Semaphore) Free() int
```
Free returns the number of available locks in the semaphore




### <a name="Semaphore.Lock">func</a> (\*Semaphore) [Lock](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1813:1839#L77)
``` go
func (s *Semaphore) Lock()
```
Lock consumes a lock in the semaphore, blocking if none is available




### <a name="Semaphore.String">func</a> (\*Semaphore) [String](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2525:2560#L106)
``` go
func (s *Semaphore) String() string
```
String returns the string representation of the semaphore




### <a name="Semaphore.Sub">func</a> (\*Semaphore) [Sub](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2241:2278#L94)
``` go
func (s *Semaphore) Sub(numLocks int)
```
Sub replaces numLocks locks in the semaphore, blocking if no locks are consumed




### <a name="Semaphore.Unlock">func</a> (\*Semaphore) [Unlock](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1939:1967#L82)
``` go
func (s *Semaphore) Unlock()
```
Unlock replaces a lock in the semaphore, blocking if no locks are consumed




### <a name="Semaphore.Until">func</a> (\*Semaphore) [Until](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1602:1641#L67)
``` go
func (s *Semaphore) Until() <-chan bool
```
Until returns a channel that fires bool(true) when the lock can be consumed.








- - -
Generated by [godoc2md](http://github.com/cognusion/godoc2md)
