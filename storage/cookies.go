package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "storage/cookies.json"

// SaveCookies saves browser cookies to a file
func SaveCookies(page *rod.Page) error {
	cookies := page.MustCookies()

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

// LoadCookies loads cookies into the browser
func LoadCookies(page *rod.Page) (bool, error) {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return false, err
	}

	var cookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return false, err
	}

	// Convert NetworkCookie → NetworkCookieParam
	var params []*proto.NetworkCookieParam
	for _, c := range cookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	page.MustSetCookies(params...)
	return true, nil
}
