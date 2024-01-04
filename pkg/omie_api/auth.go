package omie_api

import "log/slog"

// Credentials for Omie API
type Credentials struct {
	Key    string
	Secret string
}

// Set Credentials for Omie API - Key and Secret
func (c *Credentials) SetCredentials(key string, secret string) {
	c.Key = key
	c.Secret = secret
	slog.Info("Credentials setted")
}

// Config - Global Config for Omie API
var Config Credentials
