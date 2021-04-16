package routers

import (
	"github.com/gin-gonic/gin"
	"my-blog-service-go/global"
	"my-blog-service-go/pkg/limiter"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key: "/auth",
		FillInterval: time.Second,
		Capacity: 10,
		Quantum:  10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.A)
	}

}