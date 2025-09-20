package otp

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !req.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	res, err := h.svc.RequestOTP(
		c.Request.Context(),
		domain.OTPRequest{},
		req.Language,
	)
	if err != nil {
		var ae *domain.AppError
		if errors.As(err, &ae) {
			if retryableErr, ok := err.(domain.RetryableAfter); ok {
				c.Header("Retry-After", strconv.Itoa(retryableErr.RetryAfterSeconds()))
				c.JSON(MapErrCodeToHttp(ae.Code), gin.H{
					"message": ae.Message,
				})
				return
			}
			c.JSON(MapErrCodeToHttp(ae.Code), gin.H{
				"message": ae.Message,
			})
			return
		}
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: res.Status,
	})
	return
}
