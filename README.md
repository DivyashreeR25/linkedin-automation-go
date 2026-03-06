# LinkedIn Automation – Technical Proof of Concept (Go + Rod)

## Overview

This project is a **Go-based LinkedIn automation proof-of-concept** designed to demonstrate **advanced browser automation, stealth techniques, and clean system architecture**.

# Project Objectives

* Demonstrate **advanced browser automation using Go + Rod**
* Simulate **human-like behavior** to reduce bot detection signals
* Implement **robust stealth techniques** without bypassing platform security
* Build a **clean, modular, and extensible architecture**
* Showcase **intelligent automation using AI-inspired logic**

# Tech Stack

| Component          | Technology                   |
| ------------------ | ---------------------------- |
| Language           | Go (Golang)                  |
| Browser Automation | Rod                          |
| Storage            | JSON                         |
| AI Logic           | Rule-based relevance scoring |
| Platform           | Windows / Chromium           |

# Project Structure

```
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
```

# Core Functional Features

## 1️⃣ Authentication System

* Login using credentials from **environment variables**
* Detect **login success or failure**
* Persist **session cookies**
* Resume sessions **without re-login**
* Automatic **re-login if session expires**

## 2️⃣ Search & Targeting

* Search profiles using keywords
  Example: `"Software Engineer"`
* Scroll-based pagination
* Extract profile URLs
* Avoid duplicate profiles
* Configurable result limits

## 3️⃣ Profile Navigation (SAFE MODE)

* Automated navigation to profiles
* Detect presence of **"Connect" button**
* **No connection requests are sent**
* Clear logging of **connection eligibility**

## 4️⃣ Messaging System (SAFE MODE)

* Detect accepted connections via **Message button**
* Generate **personalized follow-up messages**
* Template-based message generation
* Support **dynamic variables** such as name
* Store message drafts locally in **JSON**
* **No messages are sent to real users**

# AI-Powered Profile Relevance Scoring

An **AI-inspired module** analyzes public profile content and assigns a **relevance score** for a target role (e.g., Software Engineer).

### What It Does

* Extracts visible profile text
* Applies weighted **keyword-based scoring**
* Produces **explainable relevance scores**
* Ranks profiles by suitability
* Saves results to:

```
output/profile_scores.json
```

### Why This Matters

* Demonstrates **intelligent decision-making**
* Reduces unnecessary automation actions
* Fully **ethical and read-only**
* Easily extensible to **LLMs in the future**

# Anti-Bot & Stealth Techniques (8+ Implemented)

The assignment requires **at least 8 stealth techniques**.
This project implements **more than required**.

### Mandatory Techniques

* Human-like mouse movement
* Randomized timing patterns
* Browser fingerprint masking

### Additional Techniques

* Non-headless browser execution
* Randomized viewport dimensions
* Session persistence via cookies
* Activity scheduling (business hours + breaks)
* Explicit rate limiting
* Throttling awareness
* SAFE MODE execution
* Security checkpoint detection

# Security Checkpoint Handling

The system detects:

* 2FA challenges
* CAPTCHA pages
* “Verify it's you” screens

### When Detected

* Automation **pauses safely**
* User completes verification manually
* Execution resumes **after confirmation**

⚠️ The system **does not attempt to bypass security mechanisms**.

# Ethical Design (SAFE MODE)

This project intentionally avoids:

* Sending real connection requests
* Sending real LinkedIn messages
* Bypassing LinkedIn security
* Running unattended automation

Instead, it focuses on:

* Logic demonstration
* Decision-making
* Tracking and explainability

This approach aligns with **ethical automation practices**.

# Setup Instructions

## 1️⃣ Install Go

Download and install Go:

[https://go.dev/dl/](https://go.dev/dl/)

## 2️⃣ Clone Repository

```bash
git clone https://github.com/DivyashreeR25/linkedin-automation-go.git
cd linkedin-automation
```

## 3️⃣ Configure Environment Variables

Create a `.env` file using `.env.example`.

```
LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password
```

## 4️⃣ Run the Application

```bash
go run cmd/main.go
```

# Execution Flow

```
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
```

# Output Files

```
storage/cookies.json
storage/messages.json
output/profile_scores.json
```
⚠️ **Important**

This tool runs in **SAFE MODE**, meaning it does **not perform destructive or irreversible actions** such as sending real connection requests or messages.

