package http

import (
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
}

func New(
	log *slog.Logger,
) *Handler {
	return &Handler{
		logger: log,
	}
}
