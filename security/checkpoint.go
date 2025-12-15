package security

import (
	"strings"

	"github.com/go-rod/rod"
)

func IsSecurityCheckpoint(page *rod.Page) bool {
	// 1️⃣ URL-based detection
	url := page.MustInfo().URL
	if strings.Contains(url, "checkpoint") {
		return true
	}

	// 2️⃣ HTML-based detection
	html, err := page.HTML()
	if err != nil {
		return false
	}

	signals := []string{
		"Verify",
		"verification",
		"security check",
		"captcha",
		"two-step",
		"challenge",
	}

	for _, s := range signals {
		if strings.Contains(strings.ToLower(html), strings.ToLower(s)) {
			return true
		}
	}

	return false
}
