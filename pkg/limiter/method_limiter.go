package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}


//生成新的限流对象
func NewMethodLimiter() LimiterIface  {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}
//获取限流器key值对
func (l MethodLimiter) Key(c *gin.Context) string  {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}
//获取令牌桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool)  {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

//新增多个令牌桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}