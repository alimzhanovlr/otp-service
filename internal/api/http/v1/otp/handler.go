package otp

import (
	"github.com/alimzhanovlr/otp-service/internal/domain"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	log *slog.Logger
	svc domain.UseCase
}

func NewHandler(logger *slog.Logger, svc domain.UseCase) *Handler {
	return &Handler{
		log: logger,
		svc: svc,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(MapErrCodeToHttp(domain.CodeInvalidInput), gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := h.svc.RequestOTP(
		c.Request.Context(),
		domain.OTPRequest{},
		req.Language,
	)
	if err != nil {
		if retryableErr, ok := err.(domain.RetryableAfter); ok {
			c.Header("Retry-After", strconv.Itoa(retryableErr.RetryAfterSeconds()))
			c.JSON(MapErrCodeToHttp(domain.CodeInternal), gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(MapErrCodeToHttp(domain.CodeInternal), gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: res.Status,
	})
	return
}
