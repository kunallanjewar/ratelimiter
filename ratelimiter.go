package ratelimiter

import (
	"time"
)

const (
	// TOKENS is a value for maximum tokens allowed to be
	// refilled after an interval.
	TOKENS = 2

	// INTERVAL is a default value of interval after which
	// bucket tokens are refilled.
	INTERVAL = time.Second
)

type bucket struct {
	// start can be used by a monitor go routine
	// to determine if this bucket should be deleted.
	//
	// we would delete bucket that is inactive for a while.
	start, end time.Time
	allowance  int

	// remaining tokens
	remaining int
}

func (f *bucket) expired() bool {
	return f.end.Before(time.Now().UTC())
}

// Ratelimiter implements a token
// bucket based rate limiter.
type RateLimiter struct {
	buckets map[int]*bucket
	// defaults
	dtok      int
	dinterval time.Duration
}

// New creates an instance of a Global RateLimiter.
// Global tokens are refilled after given interval.
//
// These values are applied globally and can
// be overridden per user with SetUserLimit().
func New(tokens int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		dtok:      tokens,
		dinterval: interval,
		buckets:   make(map[int]*bucket),
	}
}

// NewWithDefault return an instance of RateLimiter with default values.
//
// These values are applied globally and can be overridden per user
// with SetUserLimit().
func NewWithDefault() *RateLimiter {
	return New(TOKENS, INTERVAL)
}

// SetUserLimit overrides the global bucket
// limit for a specific user.
func (r *RateLimiter) SetUserLimit(
	id, tokens int,
	interval time.Duration) {

	b := &bucket{
		start:     time.Now().UTC(),
		end:       time.Now().UTC().Add(interval),
		allowance: tokens,
		remaining: tokens,
	}

	r.buckets[id] = b
}

// Allowed returns whether or not the request is
// allowed to be processed further.
// Returns false if token policy is violated.
func (r *RateLimiter) Allowed(user int) bool {
	// grab bucket for this user
	v, ok := r.buckets[user]
	if !ok {
		// apply global limit,
		// create new bucket but deduct one token.
		// allowance here would be default tokens.
		r.SetUserLimit(user, r.dtok-1, r.dinterval)
		return true
	}

	// bucket expired
	if v.expired() {
		// renew tokens, allow but deduct one
		v.remaining = v.allowance
		return true
	}

	if v.remaining > 0 {
		// we have tokens left
		v.remaining--
		return true
	}

	// token policy violation
	return false
}
