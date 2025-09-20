package domain

import (
	"errors"
	"fmt"
)

const (
	CodeInvalidInput    ErrorCode = "INVALID_INPUT"
	CodeNotFound        ErrorCode = "NOT_FOUND"
	CodeConflict        ErrorCode = "CONFLICT"
	CodeInternal        ErrorCode = "INTERNAL"
	CodeTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"
)

// Sentinel Errors
var (
	ErrInvalidCustomerID = errors.New("invalid customer id")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPhone      = errors.New("invalid phone")

	ErrUnknownChannel    = errors.New("unknown channel")
	ErrInvalidOTPContext = errors.New("invalid context")

	l10nErrMessages = map[ErrorCode]map[Language]string{
		CodeInternal: {
			LanguageRu: "Ошибка при отправке смс",
			LanguageEn: "Failed to send sms",
			LanguageKk: "СМС жіберу кезінде қате",
		},
	}
)

type ErrorCode string

type RetryableAfter interface {
	RetryAfterSeconds() int
}

// AppError some text
type AppError struct {
	Code    ErrorCode
	Err     error
	Message string
}

func NewAppError(code ErrorCode, err error, language Language) *AppError {
	return &AppError{
		Code:    code,
		Err:     err,
		Message: l10nErrMessages[code][language],
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %s: %s", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("%s %s", e.Code, e.Message)
}
func (e *AppError) Unwrap() error { return e.Err }

type AfterRetryableDecoratorError struct {
	*AppError
	AfterSeconds TTLSeconds
}

func (e *AfterRetryableDecoratorError) RetryAfterSeconds() int { return int(e.AfterSeconds) }

func NewAfterRetryableDecoratorError(err *AppError, afterSeconds TTLSeconds) *AfterRetryableDecoratorError {
	return &AfterRetryableDecoratorError{
		AppError:     err,
		AfterSeconds: afterSeconds,
	}
}
