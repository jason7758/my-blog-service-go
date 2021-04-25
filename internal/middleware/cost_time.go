package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CostTime()  gin.HandlerFunc  {
	return func(c *gin.Context) {
		//请求前获取当前时间
		startTime := time.Now()
		//请求处理
		c.Next()
		//处理后获取消耗时间
		endTime := time.Now()
		dur := endTime.Sub(startTime).Milliseconds()
		costTime := strconv.FormatInt(dur, 10) + "ms"
		fmt.Println("t2与t1相差：", costTime) //t2与t1相差： 50
		c.Set("cost_time", costTime)

		getCostTime, _ :=  c.Get("cost_time")
		fmt.Println("getCost：", getCostTime)
	}
}