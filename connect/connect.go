package connect

import (
	"fmt"
	"math/rand"
	"time"

	"linkedin-automation/stealth"

	"github.com/go-rod/rod"
)

// VisitProfilesSafely visits profiles WITHOUT clicking connect
// Includes explicit rate limiting & throttling awareness
func VisitProfilesSafely(page *rod.Page, profiles []string) {
	fmt.Println("\nStarting SAFE Step 9: Visiting profiles")

	// --------------------------------
	// STEP 15: Rate Limiter (explicit)
	// --------------------------------
	limiter := stealth.NewRateLimiter(5) // max 5 profiles per run

	for i, profile := range profiles {

		// Enforce rate limit
		if !limiter.Allow() {
			break
		}

		fmt.Printf("\n[%d/%d] Visiting profile: %s\n", i+1, len(profiles), profile)

		page.MustNavigate(profile)
		page.MustWaitLoad()
		time.Sleep(2 * time.Second)

		// --------------------------------
		// Human-like mouse movement
		// --------------------------------
		startX := float64(200 + rand.Intn(200))
		startY := float64(300 + rand.Intn(200))
		endX := float64(500 + rand.Intn(300))
		endY := float64(400 + rand.Intn(300))

		stealth.MoveMouseHumanLike(page, startX, startY, endX, endY)

		time.Sleep(2 * time.Second)

		// --------------------------------
		// SAFE: Detect Connect button only
		// --------------------------------
		btn, err := page.Timeout(5*time.Second).
			ElementR(`button`, `Connect`)

		if err == nil && btn != nil {
			fmt.Println("✅ Connect button detected (SAFE MODE: not clicking)")
		} else {
			fmt.Println("ℹ️ Connect button NOT found")
		}

		// --------------------------------
		// Throttling delay (human-like)
		// --------------------------------
		time.Sleep(time.Duration(3+rand.Intn(3)) * time.Second)
	}

	fmt.Println("\nSAFE Step 9 completed")
}
