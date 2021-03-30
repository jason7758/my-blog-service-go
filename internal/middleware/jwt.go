package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-blog-service-go/pkg/app"
	"my-blog-service-go/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}

			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToResponse(ecode)
			c.Abort()
			return
		}

		c.Next()

	}
}