package otp

import "github.com/alimzhanovlr/otp-service/internal/domain"

type Request struct {
	Language domain.Language
}

type SuccessResponse struct {
	Message string `json:"message"`
}
