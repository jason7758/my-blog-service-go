package middleware

import (
	"bytes"
	"my-blog-service-go/global"
	"my-blog-service-go/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)


type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields {
			"request": c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: methon: %s, status_code: %d, " + "begin_time: %d, end_time: %d"

		global.Logger.WithFields(fields).Infof(c, s, 
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}