LinkedIn Automation – Technical Proof of Concept (Go + Rod)

📌 Overview

This project is a Go-based LinkedIn automation proof-of-concept designed to demonstrate advanced browser automation, stealth techniques, and clean system architecture.

⚠️ Important:
This tool is built strictly for educational and evaluation purposes. It runs in SAFE MODE, meaning it does not perform destructive or irreversible actions such as sending real connection requests or messages.

The focus of this project is engineering capability, anti-detection awareness, and ethical automation design, not production deployment.

🎯 Project Objectives

Demonstrate advanced browser automation using Go + Rod

Simulate human-like behavior to reduce bot detection signals

Implement robust stealth techniques without bypassing platform security

Build a clean, modular, and extensible architecture

Showcase intelligent automation using AI-inspired logic

🧱 Tech Stack

Language: Go (Golang)

Browser Automation: Rod

Storage: JSON (session cookies, message drafts, AI scores)

AI Logic: Rule-based, interpretable relevance scoring (API-agnostic)

Platform: Windows / Chromium

📁 Project Structure
linkedin-automation/
├── auth/           # Login & authentication logic
├── search/         # Profile search & pagination
├── connect/        # Safe profile navigation
├── messaging/      # Messaging system (SAFE MODE)
├── security/       # Security checkpoint detection
├── stealth/        # Anti-bot & human behavior logic
├── ai/             # AI-powered relevance scoring
├── storage/        # Cookies & message storage
├── output/         # AI scoring results
├── cmd/
│   └── main.go     # Application entry point
├── .env.example    # Environment variable template
└── README.md

🔐 Core Functional Features
1️⃣ Authentication System

Login using credentials from environment variables

Detect login success/failure

Persist session cookies

Resume sessions without re-login

Automatic re-login if session becomes invalid

2️⃣ Search & Targeting

Search profiles by keyword (e.g., “Software Engineer”)

Scroll-based pagination

Extract profile URLs

Avoid duplicate profiles

Configurable result limits

3️⃣ Profile Navigation (SAFE MODE)

Programmatic navigation to profiles

Detection of “Connect” button

No connection requests are sent

Clear logging of connection eligibility

4️⃣ Messaging System (SAFE MODE)

Detect accepted connections via “Message” button

Generate personalized follow-up messages using templates

Support dynamic variables (e.g., name)

Store message drafts locally in JSON

No messages are sent to real users

🧠 AI-Powered Profile Relevance Scoring (Extra Feature)

An AI-inspired module analyzes public profile text and assigns a relevance score for a target role (e.g., Software Engineer).

What it does:

Extracts visible profile content

Applies weighted keyword-based scoring

Produces explainable relevance scores

Ranks profiles by suitability

Saves results to output/profile_scores.json

Why this matters:

Demonstrates intelligent decision-making

Reduces unnecessary automation actions

Fully ethical and read-only

Easily extensible to LLMs in the future

🛡️ Anti-Bot & Stealth Techniques (8+ Implemented)

The assignment requires at least 8 stealth techniques.
This project implements more than required.

Mandatory Techniques

Human-like mouse movement

Randomized timing patterns

Browser fingerprint masking (safe & partial)

Additional Techniques

Non-headless browser execution

Randomized viewport dimensions

Session persistence via cookies

Activity scheduling (business hours + breaks)

Explicit rate limiting per run

Throttling awareness (cooldowns & delays)

SAFE MODE (no destructive actions)

Security checkpoint detection (2FA / CAPTCHA)

🔐 Security Checkpoint Handling

The system detects:

2FA challenges

CAPTCHA pages

“Verify it’s you” screens

When detected:

Automation pauses safely

User completes verification manually

Execution resumes only after confirmation

⚠️ No attempt is made to bypass security mechanisms.

🧠 Ethical Design (SAFE MODE)

This project intentionally avoids:

Sending real connection requests

Sending real LinkedIn messages

Bypassing LinkedIn security

Running unattended background automation

Instead, it focuses on:

Logic demonstration

Decision-making

Tracking & explainability

This approach aligns with ethical automation practices.

⚙️ Setup Instructions
1️⃣ Install Go

Download from: https://go.dev/dl/

2️⃣ Clone Repository
git clone <repository-url>
cd linkedin-automation

3️⃣ Configure Environment Variables

Create a .env file using .env.example:

LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password

4️⃣ Run the Application
go run cmd/main.go

📊 Execution Flow (High-Level)
Start
↓
Activity scheduling check
↓
Launch non-headless browser
↓
Load session cookies
↓
Login verification
↓
Security checkpoint detection
↓
Search profiles
↓
Visit profiles safely
↓
Detect connect & message eligibility
↓
Generate message drafts
↓
AI relevance scoring
↓
Save outputs
↓
End

📌 Key Notes for Evaluators

This project is a technical proof-of-concept

Focus is on architecture, stealth, and correctness

SAFE MODE is intentional and ethical

AI module is explainable and API-agnostic

Designed for clarity, not abuse

👤 Author

Divyashree R
Computer Science and Engineering
Software Developer Internship Candidate