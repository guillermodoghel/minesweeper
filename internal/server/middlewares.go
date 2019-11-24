package server

import (
	"github.com/guillermodoghel/minesweeper/pkg/log"
	"github.com/guillermodoghel/minesweeper/pkg/rest"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RecoverFromPanics() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Panic("panic occurred", log.Tags{
					"panic_message": r,
					"source":        "server.RecoverFromPanics",
				})
				c.AbortWithStatusJSON(http.StatusInternalServerError, rest.NewApiError(
					"internal server error", "internal_server_error", http.StatusInternalServerError, rest.CauseList{}))
			}
		}()
		c.Next()
	}
}

func LogRequestAndResponse() func(*gin.Context) {
	return func(c *gin.Context) {
		startRequestTime := time.Now()
		c.Next()
		requestElapsedTime := time.Since(startRequestTime)
		tags := getRequestTags(c.Request)
		tags["status"] = c.Writer.Status()
		tags["response_time"] = requestElapsedTime
		if rest.IsStatusCode2XX(c.Writer.Status()) {
			log.Info("response successful", tags)
		} else {
			log.Error("response with error", nil, tags)
		}
	}
}

func getRequestTags(req *http.Request) log.Tags {
	return log.Tags{
		"request_method":       req.Method,
		"request_path":         req.URL.Path,
		log.DataKey("request"): rest.DumpRequestAsJson(req),
	}
}
