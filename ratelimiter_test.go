package ratelimiter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kunallanjewar/ratelimiter"
)

func Test_RateLimiter(t *testing.T) {
	t.Run("Global", func(t *testing.T) {
		r := ratelimiter.NewWithDefault()
		do(r, 1)
	})

	t.Run("Per User", func(t *testing.T) {
		uid := 2
		r := ratelimiter.NewWithDefault()
		r.SetUserLimit(uid, 1, ratelimiter.INTERVAL)
		do(r, uid)
	})

}

type Limiter interface {
	Allowed(int) bool
}

func do(r Limiter, id int) {
	for i := 1; i <= 10; i++ {
		n := 100
		time.Sleep(time.Duration(n) * time.Millisecond)

		if r.Allowed(id) {
			fmt.Printf("userID %v, request %v, n %v - OK\n", id, i, n)
			continue
		}

		fmt.Printf("userID %v, request %v, n %v - throttled\n", id, i, n)
	}
}
