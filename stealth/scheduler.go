package stealth

import (
	"fmt"
	"math/rand"
	"time"
)

// Check if current time is within allowed working hours
func IsWithinBusinessHours() bool {
	now := time.Now()

	hour := now.Hour()

	// Business hours: 9 AM – 6 PM
	return hour >= 9 && hour <= 18
}

// Simulate a short human break
func TakeRandomBreak() {
	breakSeconds := 5 + rand.Intn(6) // 5–10 seconds
	fmt.Printf("🛑 Taking a short break (%d seconds)...\n", breakSeconds)
	time.Sleep(time.Duration(breakSeconds) * time.Second)
}
