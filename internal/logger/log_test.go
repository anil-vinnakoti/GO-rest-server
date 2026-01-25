package logger_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/anil-vinnakoti/newsapi/internal/logger"
)

func Test_AddLoggerContextToParentContext(t *testing.T) {
	testCases := []struct {
		name           string
		ctx            context.Context
		logger         *slog.Logger
		expectedLogger bool
	}{
		{
			name: "return context without logger",
			ctx:  context.Background(),
		}, {
			name:           "return ctx as it is",
			ctx:            context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))),
			expectedLogger: true,
		}, {
			name:           "inject logger",
			ctx:            context.Background(),
			logger:         slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})),
			expectedLogger: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := logger.AddLoggerContextToParentContext(testCase.ctx, testCase.logger)

			_, ok := ctx.Value(logger.CtxKey{}).(*slog.Logger)
			if testCase.expectedLogger != ok {
				t.Errorf("expected: %v, recieved: %v", testCase.expectedLogger, ok)
			}
		})
	}
}

func Test_GetLoggerFromContext(t *testing.T) {
	testCases := []struct {
		name           string
		ctx            context.Context
		expectedLogger bool
	}{
		{
			name:           "logger exists",
			ctx:            context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))),
			expectedLogger: true,
		}, {
			name:           "new logger returned",
			ctx:            context.Background(),
			expectedLogger: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := logger.GetLoggerFromContext(testCase.ctx)

			if testCase.expectedLogger && logger == nil {
				t.Errorf("expected: %v, recieved: %v", testCase.expectedLogger, logger)
			}
		})
	}
}
