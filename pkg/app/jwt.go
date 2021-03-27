package app

import (
	"github.com/dgrijalva/jwt-go"
	"my-blog-service-go/global"
	"time"
)

type Claims struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

//获取JWT密钥
func  GetJWTSecret() []byte  {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	//nowTime := time.Now()
	//expireTime := nowTime.Add(global.JWTSetting.Expire)
	//claims := Claims{
	}
}
