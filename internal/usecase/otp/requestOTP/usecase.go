package requestOTP

import (
	"context"
	"errors"
	smsAdapter "github.com/alimzhanovlr/otp-service/internal/adapters/sms"
	"github.com/alimzhanovlr/otp-service/internal/domain"
	"github.com/alimzhanovlr/otp-service/internal/repo/otpRepo"
	"github.com/alimzhanovlr/otp-service/internal/service/otp"
	"github.com/alimzhanovlr/otp-service/internal/service/sms"
)

type usecase struct {
	Repo          domain.OTPRepository
	OTPDispatcher domain.OTPDispatcher
	SmsClient     *smsAdapter.Client
}

func New(repo domain.OTPRepository, client *smsAdapter.Client) usecase {
	return usecase{
		Repo: repo,
		OTPDispatcher: &otpDispatcher.Service{
			Repo: repo,
		},
		SmsClient: client,
	}
}

func (u *usecase) RequestOTP(ctx context.Context, req domain.OTPRequest, language string) (*domain.OTPResult, error) {
	var (
		lang       = domain.NewLanguage(language)
		retryAfter = domain.TTLSeconds(30)
	)
	if !req.Channel.IsValid() {
		return nil, domain.NewAfterRetryableDecoratorError(
			domain.NewAppError(
				domain.CodeInvalidInput,
				domain.ErrUnknownChannel,
				lang,
			), retryAfter)
	}

	if err := u.Repo.Save(ctx, req); err != nil {
		if errors.Is(err, otpRepo.ErrUniqueViolation) {
			return nil, domain.NewAppError(
				domain.CodeConflict,
				err,
				lang,
			)
		}
		return nil, domain.NewAfterRetryableDecoratorError(
			domain.NewAppError(
				domain.CodeInternal,
				err,
				lang,
			), retryAfter)
	}

	switch req.Channel {
	case domain.ChannelSMS:
		smsSender := &sms.Service{
			Client: u.SmsClient,
			SenderPayload: &sms.Payload{
				Recipient: req.Target.Value(),
				Text:      req.Context.Build(),
				Slug:      req.ID,
				Source:    "otp-service",
				Type:      1,
			},
		}
		if err := u.OTPDispatcher.DeliverAndMarkStatus(ctx, smsSender, req.ID); err != nil {
			if errors.Is(err, otpRepo.ErrNotFound) {
				return nil,
					domain.NewAppError(
						domain.CodeNotFound,
						err,
						lang,
					)
			}
			return nil,
				domain.NewAfterRetryableDecoratorError(
					domain.NewAppError(
						domain.CodeInternal,
						err,
						lang,
					), retryAfter,
				)
		}
	}

	return &domain.OTPResult{}, nil
}
