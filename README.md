# ratelimiter
--
    import "github.com/kunallanjewar/ratelimiter"

package ratelimiter is a rudimentary implementation of token bucket policy to
apply rate-limit on per user basis.

NOTE: This package is not thread-safe.

## Usage

```go
const (
	// MAX_Tokens is a value for maximum tokens allowed to be
	// refilled after an interval.
	MAX_Tokens = 10

	// MAX_Interval is a default value of interval after which
	// bucket tokens are refilled.
	MAX_Interval = time.Second
)
```

#### type RateLimiter

```go
type RateLimiter struct {
}
```

Ratelimiter implements a token bucket based rate limiter.

#### func  New

```go
func New(tokens int, duration time.Duration) *RateLimiter
```
New create an instance of RateLimiter.

#### func  NewWithDefault

```go
func NewWithDefault() *RateLimiter
```
NewWithDefault return an instance of RateLimiter with default values. These
values are applied globally and can be overridden per user with SetUserLimit().

#### func (*RateLimiter) Allowed

```go
func (r *RateLimiter) Allowed(user int) bool
```
Allowed returns whether or not the request is allowed to be processed further.
Returns false if token policy is violated.

#### func (*RateLimiter) SetUserLimit

```go
func (r *RateLimiter) SetUserLimit(id, allowance int)
```
SetUserLimit overrides the global request for a specific user.
