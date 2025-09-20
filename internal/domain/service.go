package domain

import "context"

type OTPResult struct {
	ID     int
	Status string
	TTL    TTLSeconds
}

type UseCase interface {
	RequestOTP(ctx context.Context, req OTPRequest, language Language) (OTPResult, error)
}

type Sender interface {
	Send(ctx context.Context) error
}
type OTPDispatcher interface {
	// DeliverAndMarkStatus должен еще маркировать отп как failed в случае ошибки
	DeliverAndMarkStatus(ctx context.Context, client Sender, id string) error
}
