package routers

import (
	"my-blog-service-go/global"
	"my-blog-service-go/internal/middleware"
	"my-blog-service-go/internal/routers/api"
	v1 "my-blog-service-go/internal/routers/api/v1"
	"my-blog-service-go/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwgger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	
	aritcle := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any",  ginSwgger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use() 
	{// middleware.Jwt()
		//创建标签
		apiv1.POST("/tags", tag.Create)

		//删除制定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		//更新制定标签
		apiv1.PUT("/tags/:id", tag.Update)
		//获取标签列表
		apiv1.GET("/tags", tag.List)

		//创建文章
		apiv1.POST("/articles", aritcle.Create)
		//删除指定文章
		apiv1.DELETE("/articles/:id",aritcle.Delete)
		// 更新制定文章
		apiv1.PUT("/articles/:id", aritcle.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", aritcle.Get)
		// 获取文章列表
		apiv1.GET("/articles", aritcle.List)
	}

	return r
}