package config

import "os"

type Credentials struct {
	Email    string
	Password string
}

func LoadCredentials() Credentials {
	return Credentials{
		Email:    os.Getenv("LINKEDIN_EMAIL"),
		Password: os.Getenv("LINKEDIN_PASSWORD"),
	}
}
