package ai

import (
	"strings"
)

// TEMP: simple heuristic (safe fallback if AI unavailable)
func ScoreProfile(profileURL string, profileText string) ProfileScore {
	score := 0.0

	lower := strings.ToLower(profileText)

	if strings.Contains(lower, "software engineer") {
		score += 4
	}
	if strings.Contains(lower, "backend") || strings.Contains(lower, "frontend") {
		score += 2
	}
	if strings.Contains(lower, "golang") || strings.Contains(lower, "java") || strings.Contains(lower, "python") {
		score += 2
	}

	if score > 10 {
		score = 10
	}

	reason := "Matches software engineering keywords and experience"

	return ProfileScore{
		ProfileURL: profileURL,
		Score:      score,
		Reason:     reason,
	}
}
