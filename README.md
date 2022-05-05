# ratelimiter
--
    import "github.com/kunallanjewar/ratelimiter"

ratelimiter is a rudimentary implementation of limiting based on token bucket
policy.

NOTE: This package is not thread-safe.

## Usage

```go
const (
	// MAX_Tokens is a value for maximum tokens allowed to be
	// refilled after an interval.
	TOKENS = 10

	// MAX_Interval is a default value of interval after which
	// bucket tokens are refilled.
	INTERVAL = time.Second
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
func New(tokens int, interval time.Duration) *RateLimiter
```
New creates an instance of a Global RateLimiter. Global tokens are refilled
after given interval.

These values are applied globally and can be overridden per user with
SetUserLimit().

#### func  NewWithDefault

```go
func NewWithDefault() *RateLimiter
```
NewWithDefault return an instance of RateLimiter with default values.

These values are applied globally and can be overridden per user with
SetUserLimit().

#### func (*RateLimiter) Allowed

```go
func (r *RateLimiter) Allowed(user int) bool
```
Allowed returns whether or not the request is allowed to be processed further.
Returns false if token policy is violated.

#### func (*RateLimiter) SetUserLimit

```go
func (r *RateLimiter) SetUserLimit(
	id, tokens int,
	interval time.Duration)
```
SetUserLimit overrides the global bucket limit for a specific user.
