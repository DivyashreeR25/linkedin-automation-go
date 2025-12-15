package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"linkedin-automation/ai"
	"linkedin-automation/auth"
	"linkedin-automation/config"
	"linkedin-automation/connect"
	"linkedin-automation/messaging"
	"linkedin-automation/search"
	"linkedin-automation/security"
	"linkedin-automation/stealth"
	"linkedin-automation/storage"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// Wait until LinkedIn finishes login or shows a checkpoint
func waitForLoginResult(page *rod.Page) {
	fmt.Println("⏳ Waiting for LinkedIn to complete login...")

	timeout := time.After(40 * time.Second)
	ticker := time.Tick(2 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("⚠️ Login check timeout reached")
			return
		case <-ticker:
			url := page.MustInfo().URL

			if strings.Contains(url, "/feed") {
				fmt.Println("✅ Login successful – LinkedIn feed loaded")
				return
			}

			if security.IsSecurityCheckpoint(page) {
				fmt.Println("⚠️ Security checkpoint detected during login")
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// --------------------------------
	// STEP 14: Activity Scheduling
	// --------------------------------
	if !stealth.IsWithinBusinessHours() {
		fmt.Println("⏰ Outside business hours. Exiting to mimic human behavior.")
		return
	}

	// --------------------------------
	// STEP 0: Launch browser (SAFE)
	// --------------------------------
	l := launcher.New().
		Headless(false).
		Leakless(false).
		Set("disable-blink-features", "AutomationControlled").
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	controlURL := l.MustLaunch()

	browser := rod.New().
		ControlURL(controlURL).
		MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://www.linkedin.com")

	// --------------------------------
	// STEP 11: Browser Fingerprint Masking (SAFE MODE)
	// --------------------------------
	stealth.ApplyFingerprintMask(page)

	// Randomize viewport (human-like)
	width := 1200 + rand.Intn(300)
	height := 800 + rand.Intn(200)
	page.MustSetViewport(width, height, 1, false)

	// --------------------------------
	// STEP 1: Try loading saved cookies
	// --------------------------------
	loaded, err := storage.LoadCookies(page)
	if loaded && err == nil {
		fmt.Println("✅ Session cookies loaded, checking login status...")
		page.MustReload()
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("ℹ️ No saved session, logging in...")
		creds := config.LoadCredentials()
		auth.Login(page, creds.Email, creds.Password)
	}

	// --------------------------------
	// STEP 2: Verify login
	// --------------------------------
	waitForLoginResult(page)

	// Retry login once if needed
	if !strings.Contains(page.MustInfo().URL, "/feed") {
		fmt.Println("⚠️ Session invalid, re-logging in...")
		creds := config.LoadCredentials()
		auth.Login(page, creds.Email, creds.Password)
		waitForLoginResult(page)
	}

	// --------------------------------
	// STEP 3: Security checkpoint handling
	// --------------------------------
	if security.IsSecurityCheckpoint(page) {
		log.Println("⚠️ Security checkpoint detected!")
		log.Println("➡️ Please complete verification manually in the browser.")
		log.Println("➡️ Press ENTER once done to continue...")
		fmt.Scanln()
	}

	// --------------------------------
	// STEP 4: Save cookies
	// --------------------------------
	if strings.Contains(page.MustInfo().URL, "/feed") {
		fmt.Println("💾 Saving session cookies...")
		storage.SaveCookies(page)
	}

	// --------------------------------
	// STEP 8: Search & Targeting
	// --------------------------------
	profiles := search.SearchPeople(page, "software engineer", 5)

	fmt.Println("\nCollected profile URLs:")
	for _, profile := range profiles {
		fmt.Println(profile)
	}

	// --------------------------------
	// STEP 9: SAFE Profile Visits
	// --------------------------------
	connect.VisitProfilesSafely(page, profiles)
	stealth.TakeRandomBreak()

	// --------------------------------
	// STEP 13: SAFE Messaging System
	// --------------------------------
	messaging.HandleMessagingSafely(page, profiles)

	// --------------------------------
	// STEP 10: AI PROFILE SCORING
	// --------------------------------
	fmt.Println("\n🧠 Scoring profiles using AI logic...")

	scores := []ai.ProfileScore{}

	for _, profile := range profiles {
		// VERY IMPORTANT:
		// We are NOT revisiting the profile
		// We reuse already loaded page content safely

		profileText := page.MustElement("body").MustText()

		score := ai.ScoreProfile(profile, profileText)
		scores = append(scores, score)
	}

	// STEP 7: SAVE RESULTS TO FILE
	data, _ := json.MarshalIndent(scores, "", "  ")
	os.WriteFile("output/profile_scores.json", data, 0644)

	fmt.Println("✅ AI profile scores saved to output/profile_scores.json")

	// STEP 8: SORT & DISPLAY RESULTS
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	fmt.Println("\n🏆 AI Relevance Ranking:")
	for _, s := range scores {
		fmt.Printf("Score: %.1f | %s\nReason: %s\n\n",
			s.Score,
			s.ProfileURL,
			s.Reason,
		)
	}

	fmt.Println("\nPress ENTER to close browser.")
	fmt.Scanln()
}
