package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
