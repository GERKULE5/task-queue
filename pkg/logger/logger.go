package logger

import (
	"fmt"
	"log/slog"
	"os"
	"task_queue/pkg/env"
)

func Setup() error {
	var handler slog.Handler

	mode, err := env.GetRequired("BUILD_MODE")
	if err != nil {
		return fmt.Errorf("failed to setup logger: %w", err)
	}

	switch mode {
	case "production":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05"))
				}
				return a
			},
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	slog.SetDefault(slog.New(handler))
	return nil
}
