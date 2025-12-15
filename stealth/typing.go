package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// TypeLikeHuman types text slowly with random delays
func TypeLikeHuman(element *rod.Element, text string) {
	for _, ch := range text {
		element.MustInput(string(ch))

		// random delay between keystrokes (50–150ms)
		delay := time.Duration(rand.Intn(100)+50) * time.Millisecond
		time.Sleep(delay)
	}
}
