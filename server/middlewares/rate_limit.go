package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/db"
	"github.com/hcc429/randomcat/metric"
)

var( 
	USER_EXPIRE_TIME = 5
	USER_RATE_LIMIT = 10
)
const TIME_BUCKET = 5


func RateLimit(c *gin.Context) {

	metric.IncRequests()
	userKey := getUserKey(c.ClientIP())
	// Check if user in cache
	count, err := db.GetValue(userKey)
	log.Println(count, err)
	if err != nil {
		// Redis has no User record
		// Add user to database
		log.Println("firsttime", userKey, 1)
		db.AddKeyValuePair(userKey, "1", USER_EXPIRE_TIME)
	} else {
		// Read
		visitCount, err := strconv.Atoi(count)
		if err != nil {
			log.Println("Cast to Int failed")
		} else {
			if visitCount > USER_RATE_LIMIT {
				c.Abort()
				return
			} else {
				log.Println(userKey, visitCount+1)
				db.AddKeyValuePair(userKey, strconv.Itoa(visitCount+1), USER_EXPIRE_TIME)
			}
		}
	}
	c.Next() // forward to controller
	if c.Writer.Status() < 400 {
		metric.IncSuccessfulRequests()
	}
}

func getUserKey(IP string) string{
	bucket := time.Now().Unix() / TIME_BUCKET
	IP = IP + "_" + strconv.FormatInt(bucket, 10)
	return IP
}
