package otp

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/", h.Handle)
}
