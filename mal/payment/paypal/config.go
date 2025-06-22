package paypal

import (
	"fmt"
	"os"
)

type Config struct {
	ClientID     string
	ClientSecret string
	WebhookID    string
	Mode         string // "sandbox" or "live"
}

func LoadConfig() (*Config, error) {
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	if clientID == "" {
		return nil, fmt.Errorf("invalid client id")
	}

	clientSecret := os.Getenv("PAYPAL_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, fmt.Errorf("invalid client secret")
	}

	webhookID := os.Getenv("PAYPAL_WEBHOOK_ID")
	if webhookID == "" {
		return nil, fmt.Errorf("invalid webhook id")
	}

	mode := os.Getenv("PAYPAL_MODE")
	if mode == "" {
		mode = "sandbox"
	}

	return &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		WebhookID:    webhookID,
		Mode:         mode,
	}, nil
}

func (c *Config) Sandbox() bool {
	return c.Mode == "sandbox"
}

func (c *Config) BaseURL() string {
	if c.Sandbox() {
		return "https://api-m.sandbox.paypal.com"
	}
	return "https://api-m.paypal.com"
}

func (c *Config) AuthURL() string {
	if c.Sandbox() {
		return "https://www.sandbox.paypal.com"
	}
	return "https://www.paypal.com"
}

func (c *Config) AccessTokenKey() string {
	return "paypal:access_token:"
}
