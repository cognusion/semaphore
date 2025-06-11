

# semaphore
`import "github.com/cognusion/semaphore"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package semaphore is a super simple goro-safe semaphore struct for Go.


##### Example :
``` go
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
fmt.Printf("Free locks? %d\n", S.Free())
time.Sleep(3 * time.Second)
fmt.Printf("Free locks now? %d\n", S.Free())
```



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
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
  * [func (s *Semaphore) IsFree(freeFor time.Duration) &lt;-chan bool](#Semaphore.IsFree)
  * [func (s *Semaphore) Lock()](#Semaphore.Lock)
  * [func (s *Semaphore) String() string](#Semaphore.String)
  * [func (s *Semaphore) Sub(numLocks int)](#Semaphore.Sub)
  * [func (s *Semaphore) Unlock()](#Semaphore.Unlock)
  * [func (s *Semaphore) Until() &lt;-chan bool](#Semaphore.Until)

#### <a name="pkg-examples">Examples</a>
* [Package](#example-)

#### <a name="pkg-files">Package files</a>
[counter.go](https://github.com/cognusion/semaphore/tree/master/counter.go) [numlock.go](https://github.com/cognusion/semaphore/tree/master/numlock.go) [semaphore.go](https://github.com/cognusion/semaphore/tree/master/semaphore.go)


## <a name="pkg-constants">Constants</a>
``` go
const UntilFreeTimeout = 2 * time.Millisecond
```
UntilFreeTimeout is a Duration for the goro spawned by the Until func to wait
if there is nothing consuming messages from provided chan, before releasing
the lock.





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




## <a name="Semaphore">type</a> [Semaphore](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=385:426#L15)
``` go
type Semaphore struct {
    // contains filtered or unexported fields
}

```
Semaphore is a goro-safe simple semaphore







### <a name="NewSemaphore">func</a> [NewSemaphore](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=508:545#L20)
``` go
func NewSemaphore(size int) Semaphore
```
NewSemaphore returns a Semaphore allowing up to 'size' locks before blocking





### <a name="Semaphore.Add">func</a> (\*Semaphore) [Add](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1568:1605#L58)
``` go
func (s *Semaphore) Add(numLocks int)
```
Add consumes numLocks locks in the semaphore, blocking if none is available




### <a name="Semaphore.Free">func</a> (\*Semaphore) [Free](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1900:1930#L72)
``` go
func (s *Semaphore) Free() int
```
Free returns the number of available locks in the semaphore




### <a name="Semaphore.IsFree">func</a> (\*Semaphore) [IsFree](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2109:2170#L78)
``` go
func (s *Semaphore) IsFree(freeFor time.Duration) <-chan bool
```
IsFree takes a Duration, and makes a decent try on determining if someone consumed
a lock over the Duration, ala a WaitGroup.Wait().




### <a name="Semaphore.Lock">func</a> (\*Semaphore) [Lock](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1319:1345#L48)
``` go
func (s *Semaphore) Lock()
```
Lock consumes a lock in the semaphore, blocking if none is available




### <a name="Semaphore.String">func</a> (\*Semaphore) [String](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=2612:2647#L101)
``` go
func (s *Semaphore) String() string
```
String returns the string representation of the semaphore




### <a name="Semaphore.Sub">func</a> (\*Semaphore) [Sub](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1747:1784#L65)
``` go
func (s *Semaphore) Sub(numLocks int)
```
Sub replaces numLocks locks in the semaphore, blocking if no locks are consumed




### <a name="Semaphore.Unlock">func</a> (\*Semaphore) [Unlock](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=1445:1473#L53)
``` go
func (s *Semaphore) Unlock()
```
Unlock replaces a lock in the semaphore, blocking if no locks are consumed




### <a name="Semaphore.Until">func</a> (\*Semaphore) [Until](https://github.com/cognusion/semaphore/tree/master/semaphore.go?s=854:893#L29)
``` go
func (s *Semaphore) Until() <-chan bool
```
Until returns a channel that fires bool(true) when the lock can be consumed.
If nothing is listening to the returned channel, after 2ms (UntilFreeTimeout)
the lock will be removed, so if you're not paying attention, don't use this function.








- - -
Generated by [godoc2md](http://github.com/cognusion/godoc2md)
