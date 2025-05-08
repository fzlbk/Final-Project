package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		c.Set("RequestID", reqID)
		start := time.Now()
		c.Next()
		dur := time.Since(start)
		log.Printf("[%s] [RequestID: %s] %s %s - %d - Duration: %.3fms",
			time.Now().UTC().Format(time.RFC3339),
			reqID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			float64(dur.Microseconds())/1000,
		)
	}
}
