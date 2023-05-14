package controller

import (
	"ginStudy/common"
	"ginStudy/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexPage(c *gin.Context) {
	key := c.DefaultQuery("key", "key")
	isKey := key != ""

	value := model.GetKey(key)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "",
		"option":  common.OptionMap,
		"value":   value,
		"isKey":   isKey,
	})
}

func Get404Page(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"message": "",
		"option":  common.OptionMap,
	})
}
