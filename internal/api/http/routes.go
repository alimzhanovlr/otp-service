package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SetupRoutes(dir string) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	return router.Handler()
}
