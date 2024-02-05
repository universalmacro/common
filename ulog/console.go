package ulog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func NewHandler(level slog.Level) *ConsoleHandler {
	return &ConsoleHandler{
		enabled: level,
	}
}

type ConsoleHandler struct {
	enabled slog.Level
	attr    []slog.Attr
}

func (h *ConsoleHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.enabled <= level
}

func (h *ConsoleHandler) Handle(ctx context.Context, r slog.Record) error {
	var fields []string
	for _, attr := range h.attr {
		fields = append(fields, fmt.Sprintf("%s=%v", attr.Key, attr.Value))
	}
	r.Attrs(func(a slog.Attr) bool {
		fields = append(fields, fmt.Sprintf("%s=%v", a.Key, a.Value))
		return true
	})
	fmt.Println(strings.Join(fields, " "))
	return nil
}

func (h ConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attr = append(h.attr, attrs...)
	return &h
}

func (h ConsoleHandler) WithGroup(name string) slog.Handler {
	return &h
}
