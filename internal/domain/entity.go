package domain

import (
	"time"
)

type OTPRequest struct {
	ID           string
	Target       Target
	Channel      Channel
	Context      OTPContext
	TTLSeconds   TTLSeconds
	TemplateKey  string
	CodeHash     string
	ExpiresAt    time.Time
	VerifiedAt   *time.Time
	Attempts     int
	Status       Status
	TemplateData map[string]any
}
