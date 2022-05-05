# ratelimiter
--
    import "github.com/kunallanjewar/ratelimiter"

ratelimiter is a proof-of-concept implementation of request limiting based on
token bucket approach with a caveat that rate of refill is equal to the bucket
capacity after each interval.

🚫 This package is neither thread-safe nor intended to be used in any production
environment.

### Needed Improvements

    1. Allow refill at a lower rate per interval until bucket cap is reached.
    2. Implement thread safety.
    3. Allow for some amount of traffic burst.

## Usage

```go
const (
	// TOKENS is a value for maximum tokens allowed to be
	// refilled after an interval.
	TOKENS = 2

	// INTERVAL is a default value of interval after which
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
