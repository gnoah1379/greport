package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		var level zerolog.Level
		switch {
		case status >= http.StatusBadRequest:
			level = zerolog.ErrorLevel
		default:
			level = zerolog.InfoLevel
		}
		log.WithLevel(level).
			Int("status", status).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Dur("latency", latency).
			Send()
	}
}
