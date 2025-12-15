package stealth

import "fmt"

type RateLimiter struct {
	MaxProfiles int
	Visited     int
}

func NewRateLimiter(max int) *RateLimiter {
	return &RateLimiter{
		MaxProfiles: max,
		Visited:     0,
	}
}

func (r *RateLimiter) Allow() bool {
	if r.Visited >= r.MaxProfiles {
		fmt.Println("⛔ Rate limit reached. Stopping further actions.")
		return false
	}
	r.Visited++
	return true
}
