package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"my-blog-service-go/global"
)

func Tracing() func(c *gin.Context)  {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header)
			)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(), global.Tracer, c.Request.URL.Path)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				//c.Request.Context()
			)
		}
	}
}