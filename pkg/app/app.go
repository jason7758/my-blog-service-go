package app

import (
	"github.com/gin-gonic/gin"
	"my-blog-service-go/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type pager struct {
	//页码
	Page int `json:"page"`
	//每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

//生成响应数据
func (r *Response) ToResponse(data interface{})  {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//返回带有错误的响应数据
func (r *Response) ToErrorResponse(err *errcode.Error)  {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0  {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}



