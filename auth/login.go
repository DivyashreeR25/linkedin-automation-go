package auth

import (
	"fmt"
	"strings"
	"time"

	"linkedin-automation/stealth"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page, email string, password string) {
	fmt.Println("Opening LinkedIn login page")

	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	// Email
	emailInput := page.MustElement("input#username")
	stealth.TypeLikeHuman(emailInput, email)

	time.Sleep(1 * time.Second)

	// Password
	passwordInput := page.MustElement("input#password")
	stealth.TypeLikeHuman(passwordInput, password)

	time.Sleep(1 * time.Second)

	// Submit
	signInBtn := page.MustElement("button[type='submit']")
	signInBtn.MustClick()

	fmt.Println("Login submitted, checking status...")
	time.Sleep(5 * time.Second)

	checkLoginStatus(page)
}

func checkLoginStatus(page *rod.Page) {
	url := page.MustInfo().URL
	url = strings.ToLower(url)

	switch {
	case strings.Contains(url, "feed"):
		fmt.Println("✅ Login successful – LinkedIn feed loaded")

	case strings.Contains(url, "checkpoint"):
		fmt.Println("⚠️ Security checkpoint / verification detected")
		fmt.Println("Automation halted for safety")

	case strings.Contains(url, "challenge"):
		fmt.Println("⚠️ Captcha or challenge detected")
		fmt.Println("Automation halted for safety")

	default:
		fmt.Println("❌ Login status unclear – manual review required")
	}
}
