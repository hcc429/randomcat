package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/metric"
)



func RateLimit(c *gin.Context){
	// TODO: rate limit
	metric.IncRequests()
	c.Next()
	if c.Writer.Status() < 400{
		metric.IncSuccessfulRequests()
	}
}