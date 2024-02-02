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
	attr    []slog.Attr
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.enabled <= level
}

func (h *Handler) Handle(ctx context.Context, e slog.Record) error {
	fmt.Println(e.Message)
	for _, attr := range h.attr {
		fmt.Printf("%s=%v\n", attr.Key, attr.Value)
	}
	return nil
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attr = append(h.attr, attrs...)
	return &h
}

func (h Handler) WithGroup(name string) slog.Handler {
	return &h
}
