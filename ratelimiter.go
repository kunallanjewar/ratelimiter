package ratelimiter

import (
	"time"
)

const (
	// MAX_Tokens is a value for maximum tokens allowed to be
	// refilled after an interval.
	MAX_Tokens = 10

	// MAX_Interval is a default value of interval after which
	// bucket tokens are refilled.
	MAX_Interval = time.Second
)

type bucket struct {
	// start can be used by a monitor go routine
	// to determine if this bucket should be deleted.
	//
	// we would delete bucket that is inactive for a while.
	start, end time.Time
	interval   time.Duration
	allowance  int
	tokens     int
}

func (f *bucket) expired() bool {
	return f.end.Before(time.Now())
}

// Ratelimiter implements a token bucket based rate limiter.
type RateLimiter struct {
	bucket map[int]*bucket
	dtok   int
}

// New create an instance of RateLimiter.
func New(tokens int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		dtok:   tokens,
		bucket: make(map[int]*bucket),
	}
}

// NewWithDefault return an instance of RateLimiter with default values.
// These values are applied globally and can be overridden per user with SetUserLimit().
func NewWithDefault() *RateLimiter {
	return New(MAX_Tokens, MAX_Interval)
}

// SetUserLimit overrides the global request for a specific user.
func (r *RateLimiter) SetUserLimit(id, allowance int) {
	b := &bucket{
		start:     time.Now(),
		end:       time.Now().Add(time.Second),
		interval:  time.Second,
		allowance: allowance,
		tokens:    allowance,
	}

	r.bucket[id] = b
}

// Allowed returns whether or not the request is allowed to be processed further.
// Returns false if token policy is violated.
func (r *RateLimiter) Allowed(user int) bool {
	// grab bucket for this user
	v, ok := r.bucket[user]
	if !ok {
		// apply global limit, create new bucket but deduct one token
		// allowance here would be default tokens.
		r.SetUserLimit(user, r.dtok-1)
		return true
	}

	// bucket expired
	if v.expired() {
		// renew tokens, allow but deduct one
		v.tokens = v.allowance
		return true
	}

	if v.tokens > 0 {
		// we have tokens left
		v.tokens--
		return true
	}

	// exceeded token limit within window
	return false
}
