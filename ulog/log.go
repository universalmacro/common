package ulog

import (
	"context"
	"fmt"
	"log/slog"
)

func NewHandler(level slog.Level) *Handler {
	return &Handler{
		enabled: level,
	}
}

type Handler struct {
	enabled slog.Level
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.enabled <= level
}

func (h *Handler) Handle(ctx context.Context, e *slog.Record) error {
	if e == nil {
		return nil
	}
	fmt.Println(*e)
	return nil
}

func (h *Handler) WithAttrs(attrs []slog.Attr) Handler {
	return *h
}

func (h *Handler) WithGroup(name string) Handler {
	return *h
}
