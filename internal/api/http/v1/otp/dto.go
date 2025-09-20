package otp

import "github.com/alimzhanovlr/otp-service/internal/domain"

type Request struct {
	Language domain.Language
}

func (r Request) Validate() bool {
	return true
}

type SuccessResponse struct {
	Message string `json:"message"`
}
