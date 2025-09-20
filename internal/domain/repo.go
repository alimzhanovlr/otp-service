package domain

import "context"

type OTPRepository interface {
	Save(ctx context.Context, req OTPRequest) error
	Get(ctx context.Context, id string) (OTPRequest, error)
	MarkOTPStatus(ctx context.Context, id string, status Status) error
}
