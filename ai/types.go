package ai

type ProfileScore struct {
	ProfileURL string  `json:"profile_url"`
	Score      float64 `json:"score"`
	Reason     string  `json:"reason"`
}
