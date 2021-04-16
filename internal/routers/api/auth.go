package api

import (
	"github.com/gin-gonic/gin"
	"my-blog-service-go/global"
	"my-blog-service-go/internal/service"
	"my-blog-service-go/pkg/app"
	"my-blog-service-go/pkg/errcode"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Error(c, "app.BindAndVaild errs:%v", errs)
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)

	if err != nil {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs )
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}