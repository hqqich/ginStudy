package middleware

import (
	"context"
	"fmt"
	"ginStudy/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var timeFormat = "2006-01-02T15:04:05.000Z"

// 使用gin解决跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")  //允许访问的域名 *表示所有的都可以访问
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")   //缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*") //可以通过访问的方法 get post ...  *表示允许所有方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*") //允许请求带的header头信息
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

func rateLimitHelper(c *gin.Context, maxRequestPerMinute int, mark string) {
	ctx := context.Background()
	rdb := common.RDB
	key := "rateLimit:" + mark + c.ClientIP()
	listLength, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	if listLength < int64(maxRequestPerMinute) {
		rdb.LPush(ctx, key, time.Now().Format(timeFormat))
	} else {
		oldTimeStr, _ := rdb.LIndex(ctx, key, -1).Result()
		oldTime, err := time.Parse(timeFormat, oldTimeStr)
		if err != nil {
			fmt.Println(err)
		}
		newTimeStr := time.Now().Format(timeFormat)
		newTime, err := time.Parse(timeFormat, newTimeStr)
		if err != nil {
			fmt.Println(err)
		}
		// time.Since will return negative number!
		// See: https://stackoverflow.com/questions/50970900/why-is-time-since-returning-negative-durations-on-windows
		if newTime.Sub(oldTime).Seconds() < 60 {
			c.Status(http.StatusTooManyRequests)
			c.Abort()
			return
		} else {
			rdb.LPush(ctx, key, time.Now().Format(timeFormat))
			rdb.LTrim(ctx, key, 0, int64(maxRequestPerMinute-1))
		}
	}
}

func GlobalWebRateLimit() func(c *gin.Context) {
	return func(c *gin.Context) {
		if common.RedisEnabled {
			rateLimitHelper(c, common.GlobalWebRateLimit, "GW")
		} else {
			c.Next()
		}
	}
}

func GlobalAPIRateLimit() func(c *gin.Context) {
	return func(c *gin.Context) {
		if common.RedisEnabled {
			rateLimitHelper(c, common.GlobalApiRateLimit, "GA")
		} else {
			c.Next()
		}
	}
}

func CriticalRateLimit() func(c *gin.Context) {
	return func(c *gin.Context) {
		if common.RedisEnabled {
			rateLimitHelper(c, common.CriticalRateLimit, "CT")
		} else {
			c.Next()
		}

	}
}

func DownloadRateLimit() func(c *gin.Context) {
	return func(c *gin.Context) {
		if common.RedisEnabled {
			rateLimitHelper(c, common.DownloadRateLimit, "CM")
		} else {
			c.Next()
		}
	}
}
