package ai

func BuildPrompt(profileText string, role string) string {
	return `
You are an expert technical recruiter.

Analyze the following LinkedIn profile content and score its relevance
for the role of "` + role + `" on a scale of 1 to 10.

Rules:
- Consider skills, job titles, experience
- Ignore formatting or missing sections
- Be concise

Respond in this format:
Score: <number>
Reason: <short explanation>

Profile Content:
` + profileText
}
