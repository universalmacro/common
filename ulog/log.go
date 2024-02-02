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

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Println(r.Message)
	for _, l := range h.attr {
		fmt.Printf("%s=%v ", l.Key, l.Value)
	}
	r.Attrs(func(a slog.Attr) bool {
		fmt.Printf("%s=%v ", a.Key, a.Value)
		return true
	})
	return nil
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attr = append(h.attr, attrs...)
	return &h
}

func (h Handler) WithGroup(name string) slog.Handler {
	return &h
}
