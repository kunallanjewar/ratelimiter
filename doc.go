// ratelimiter is a proof-of-concept implementation of
// request limiting based on token bucket approach with
// a caveat that rate of refill is equal to the bucket
// capacity after each interval.
//
// 🚫 This package is neither thread-safe nor intended to be used in any
// production environment.
//
// Needed Improvements
//	1. Allow refill at a lower rate per interval until bucket cap is reached.
//	2. Implement thread safety.
//	3. Allow for some amount of traffic burst.
package ratelimiter
