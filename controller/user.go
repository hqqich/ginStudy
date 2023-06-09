package controller

import (
	"context"
	"encoding/json"
	"ginStudy/common"
	"ginStudy/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	validate := user.Validate() // 验证是否包含特殊字符
	user.ValidateAndFill()
	if user.Status != common.UserStatusEnabled || !validate {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "用户名或密码错误，或者该用户已被封禁",
			"option":  common.OptionMap,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("username", username)
	session.Set("role", user.Role)
	err := session.Save()
	if err != nil {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "无法保存会话信息，请重试",
			"option":  common.OptionMap,
		})
		return
	}
	redirectUrl := c.Request.Referer()
	if strings.HasSuffix(redirectUrl, "/login") {
		redirectUrl = "/"
	}
	c.Redirect(http.StatusFound, redirectUrl)
	return
}

func LoginByToken(c *gin.Context) {

	//使用map获取请求参数
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)

	////使用结构体承接body参数
	//var requestUser = model.User{}
	//json.NewDecoder(c.Request.Body).Decode(&requestUser)
	//
	//
	////gin框架提供的绑定参数
	//var requestUser = model.User{}
	//c.Bind(&requestUser)

	user := model.User{
		Username: requestMap["username"],
		Password: requestMap["password"],
	}

	validate := user.Validate() // 验证是否包含特殊字符
	user.ValidateAndFill()
	if user.Status != common.UserStatusEnabled || !validate {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "用户名或密码错误，或者该用户已被封禁",
			"option":  common.OptionMap,
		})
		return
	}

	// 生成token,写到redis
	jwt := model.GetJwt(user)

	ctx := context.Background()
	rdb := common.RDB

	marshal, _ := json.Marshal(user)
	err := rdb.Set(ctx, "token:"+jwt, marshal, time.Duration(60*time.Second)).Err()
	if err != nil {
		panic(err)
	}

	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}

	c.JSON(200, gin.H{
		"code": 1,
		"data": jwt,
	})
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}

func UpdateSelf(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	user.Id = c.GetInt("id")
	role := c.GetInt("role")
	if role != common.RoleAdminUser {
		user.Role = 0
		user.Status = 0
	}
	// TODO: check Display Name to avoid XSS attack
	if err := user.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

// CreateUser Only admin user can call this, so we can trust it
func CreateUser(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	user.DisplayName = user.Username
	// TODO: Check user.Status && user.Role
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	if err := user.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

type ManageRequest struct {
	Username string `json:"username"`
	Action   string `json:"action"`
}

// ManageUser Only admin user can do this
func ManageUser(c *gin.Context) {
	var req ManageRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	user := model.User{
		Username: req.Username,
	}
	// Fill attributes
	model.DB.Where(&user).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
		})
		return
	}
	switch req.Action {
	case "disable":
		user.Status = common.UserStatusDisabled
	case "enable":
		user.Status = common.UserStatusEnabled
	case "delete":
		if err := user.Delete(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	case "promote":
		user.Role = common.RoleAdminUser
	case "demote":
		user.Role = common.RoleCommonUser
	}

	if err := user.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
