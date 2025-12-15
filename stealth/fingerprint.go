package stealth

import "github.com/go-rod/rod"

// ApplyFingerprintMask applies SAFE browser fingerprint masking
func ApplyFingerprintMask(page *rod.Page) {
	page.MustEval(`() => {
		// navigator.webdriver = undefined
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		// Fake platform
		Object.defineProperty(navigator, 'platform', {
			get: () => 'Win32'
		});

		// Fake languages
		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});

		// Fake hardwareConcurrency
		Object.defineProperty(navigator, 'hardwareConcurrency', {
			get: () => 8
		});

		// Fake deviceMemory
		Object.defineProperty(navigator, 'deviceMemory', {
			get: () => 8
		});

		// Override userAgent
		Object.defineProperty(navigator, 'userAgent', {
			get: () => 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36'
		});

		return true;
	}`)
}
