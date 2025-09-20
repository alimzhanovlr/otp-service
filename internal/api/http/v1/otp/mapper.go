package otp

import (
	"github.com/alimzhanovlr/otp-service/internal/domain"
	"net/http"
)

func MapErrCodeToHttp(code domain.ErrorCode) int {
	switch code {
	case domain.CodeInternal:
		return http.StatusInternalServerError
	case domain.CodeInvalidInput:
		return http.StatusUnprocessableEntity
	case domain.CodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusOK
	}
}
