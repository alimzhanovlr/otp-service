package sms

import (
	"context"
	"github.com/alimzhanovlr/otp-service/internal/adapters/sms"
)

type Service struct {
	Client        *sms.Client
	SenderPayload *Payload
}

func (s *Service) Send(ctx context.Context) error {
	return s.Client.Send(ctx, sms.Payload{
		Recipient: s.SenderPayload.Recipient,
		Text:      s.SenderPayload.Text,
		Slug:      s.SenderPayload.Slug,
		Source:    s.SenderPayload.Source,
		Type:      s.SenderPayload.Type,
	})
}
