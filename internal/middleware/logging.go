package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type LoggingMiddleware struct {
	Logger *slog.Logger
}

func NewLoggingMiddleware(logger *slog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{Logger: logger}
}

func (lm *LoggingMiddleware) Log(c *gin.Context) {
	lm.Logger.Info(
		"Request",
		"method", c.Request.Method,
		"path", c.FullPath(),
		"ip", c.ClientIP(),
	)
	c.Next()
}
