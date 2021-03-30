package middleware

import (
	"github.com/gin-gonic/gin"
	"my-blog-service-go/pkg/app"
	"my-blog-service-go/pkg/errcode"
	"my-blog-service-go/pkg/limiter"
)

//进行限流处理，如果获取Bucket未成功，则返回请求太多的响应信息
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}

}
