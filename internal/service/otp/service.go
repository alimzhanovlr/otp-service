package otp

import (
	"context"
	"github.com/alimzhanovlr/otp-service/internal/domain"
)

type Service struct {
	Repo domain.OTPRepository
}

func (s *Service) DeliverAndMarkStatus(ctx context.Context, client domain.Sender, id string) error {
	if err := client.Send(ctx); err != nil {
		s.Repo.MarkOTPStatus(ctx, id, domain.StatusFailed)
		return err
	}
	s.Repo.MarkOTPStatus(ctx, id, domain.StatusSent)
	return nil
}
