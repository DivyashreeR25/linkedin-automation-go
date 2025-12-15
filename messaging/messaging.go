package messaging

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

type MessageLog struct {
	ProfileURL string    `json:"profile_url"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

const messageFile = "storage/messages.json"

// Detect if Message button exists
func CanMessage(page *rod.Page) bool {
	btn, err := page.Timeout(5*time.Second).
		ElementR("button", "Message")
	return err == nil && btn != nil
}

// Generate personalized message (template-based)
func GenerateMessage(profileURL string) string {
	name := extractName(profileURL)

	template := "Hi {{name}}, I came across your profile and would love to connect and learn more about your work."

	return strings.Replace(template, "{{name}}", name, 1)
}

// Save message log to JSON
func SaveMessage(profileURL, message string) error {
	var logs []MessageLog

	if data, err := os.ReadFile(messageFile); err == nil {
		_ = json.Unmarshal(data, &logs)
	}

	logs = append(logs, MessageLog{
		ProfileURL: profileURL,
		Message:    message,
		Timestamp:  time.Now(),
	})

	data, _ := json.MarshalIndent(logs, "", "  ")
	return os.WriteFile(messageFile, data, 0644)
}

// Very simple name extractor from URL
func extractName(url string) string {
	parts := strings.Split(url, "/in/")
	if len(parts) < 2 {
		return "there"
	}
	name := strings.Split(parts[1], "-")[0]
	return strings.Title(name)
}

// SAFE MODE messaging (NO SEND)
func HandleMessagingSafely(page *rod.Page, profiles []string) {
	fmt.Println("\n📨 STEP 13: Messaging System (SAFE MODE)")

	for _, profile := range profiles {
		page.MustNavigate(profile)
		page.MustWaitLoad()
		time.Sleep(3 * time.Second)

		if CanMessage(page) {
			msg := GenerateMessage(profile)
			_ = SaveMessage(profile, msg)

			fmt.Println("🟢 Message prepared (NOT sent):", profile)
			fmt.Println("   ➤", msg)
		} else {
			fmt.Println("⚪ Cannot message (not connected):", profile)
		}
	}
}
