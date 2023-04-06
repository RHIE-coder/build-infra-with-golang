# context

## [Core Concept]

### - Structure

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

### - Getting

 - `func Background() Context`
 - `func TODO() Context`

### - Setting

 - `func WithValue(parent Context, key, val any) Context`
 - `func WithCancel(parent Context) (ctx Context, cancel CancelFunc)`
 - `func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)`
 - `func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)`
 - `func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)`

### - Working
 - `func Cause(c Context)`
 - `type CancelCauseFunc func(cause error)`
 - `type CancelFunc func()`