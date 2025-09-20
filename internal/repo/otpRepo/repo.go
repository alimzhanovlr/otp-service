package otpRepo

import (
	"context"
	"github.com/alimzhanovlr/otp-service/internal/domain"
)

type Repo struct {
	*pgxpool.Pool
}

func (r Repo) Save(ctx context.Context, req domain.OTPRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r Repo) Get(ctx context.Context, id string) (domain.OTPRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repo) MarkOTPStatus(ctx context.Context, id string, status domain.Status) error {
	//TODO implement me
	panic("implement me")
}
