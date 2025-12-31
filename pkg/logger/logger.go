package logger

import (
	"context"

	"go.uber.org/zap"
)

const loggerKey = "logger"

type Logger struct {
	l zap.Logger
}

func NewLogger(env string) (*Logger, error) {
	var zapLogger *zap.Logger
	if env != "dev" {
		zapLogger, _ = zap.NewProduction()
	} else {
		zapLogger, _ = zap.NewDevelopment()
	}
	return &Logger{l: *zapLogger}, nil
}
func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}
func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}
func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}
func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)
}
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(loggerKey).(*Logger); ok {
		return logger
	}
	return nil
}
