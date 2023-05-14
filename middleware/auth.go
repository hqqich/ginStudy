package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"ginStudy/common"
	"ginStudy/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{
				"message": "未登录或登录已过期",
				"option":  common.OptionMap,
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", session.Get("role"))
		c.Set("id", session.Get("id"))
		c.Next()
	}
}

func ApiAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，请登录后重试",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", session.Get("role"))
		c.Set("id", session.Get("id"))
		c.Next()
	}
}

func ApiAdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get("role")
		if role == nil || role != common.RoleAdminUser {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，请检查你是否登录或者有相关权限",
			})
			c.Abort()
			return
		}
		c.Set("username", session.Get("username"))
		c.Set("role", role)
		c.Set("id", session.Get("id"))
		c.Next()
	}
}

func TokenAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 从请求头中拿数据
		token := c.Request.Header.Get("Authorization")
		// Do something with the token, like validate it
		if token == "" {
			//c.AbortWithStatus(http.StatusUnauthorized)
			//return
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作",
			})
			c.Abort()
			return
		}

		ctx := context.Background()
		rdb := common.RDB

		// 获取redis中数据
		val, err := rdb.Get(ctx, "token:"+token).Result()
		if err != nil {
			panic(err)
		}

		if val == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作",
			})
			c.Abort()
			return
		}

		user := model.User{}

		json.Unmarshal([]byte(val), &user)

		fmt.Println(user)

		//session := sessions.Default(c)
		//username := session.Get("username")
		//if username == nil {
		//	c.JSON(http.StatusForbidden, gin.H{
		//		"success": false,
		//		"message": "无权进行此操作，请登录后重试",
		//	})
		//	c.Abort()
		//	return
		//}
		//c.Set("username", username)
		//c.Set("role", session.Get("role"))
		//c.Set("id", session.Get("id"))

		c.Next()
	}
}
