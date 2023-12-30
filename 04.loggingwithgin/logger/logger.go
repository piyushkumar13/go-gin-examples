package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func FormatLogs(param gin.LogFormatterParams) string {

	return fmt.Sprintf(" %s %s %s %s %s %d %s %s %s",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339),
		param.Method,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)

}
