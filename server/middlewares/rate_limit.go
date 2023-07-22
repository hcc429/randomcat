package middlewares

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/cache"
	"github.com/hcc429/randomcat/metrics"
)

const (
	RATE_LIMIT_INTERVAL = 10
	RATE_LIMIT_QUOTA    = 5
)

func RateLimit(c *gin.Context) {

	userKey := getUserKey(c.ClientIP())
	count, err := cache.GetValue(userKey)
	if err != nil {
		// Redis has no User record
		// Add user to database
		cache.AddKeyValuePair(userKey, "1", RATE_LIMIT_INTERVAL)
	} else {
		// Read
		visitCount, err := strconv.Atoi(count)
		if err != nil {
			log.Println("Cast to Int failed")
		} else {
			visitCount++
			if visitCount > RATE_LIMIT_QUOTA {
				c.AbortWithError(429, errors.New("Too Many Request!"))
				return
			} else {
				cache.AddKeyValuePair(userKey, visitCount, RATE_LIMIT_INTERVAL)
			}
		}
		log.Println(userKey, visitCount)
	}
	c.Next() // forward to controller
	if c.Writer.Status() < 400 {
		metric.IncSuccessfulRequests()
	} else {
		metric.IncFailedRequests()
	}
}

func getUserKey(IP string) string {
	bucket := time.Now().Unix() / int64(RATE_LIMIT_INTERVAL)
	IP = IP + "_" + strconv.FormatInt(bucket, 10)
	return IP
}
