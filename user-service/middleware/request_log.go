package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
	"user-service/infrastructure/log"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := uuid.New().String()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		ctx := context.WithValue(timeoutCtx, "reqId", reqId)
		c.Request = c.Request.WithContext(ctx)

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)

		requestLog := logrus.Fields{
			"request_id": reqId,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"latency":    latency,
		}

		if c.Writer.Status() == 200 || c.Writer.Status() == 201 {
			log.Logger.WithFields(requestLog).Info("Request Success")
		} else {
			log.Logger.WithFields(requestLog).Error("Request Failed")
		}

	}
}
