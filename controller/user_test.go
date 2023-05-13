package controller

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

// jwt测试
func TestHelloWorld(t *testing.T) {

	// 创建秘钥
	key := []byte("aaa")
	// 创建Token结构体
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "zhangshan",
		"pass": "123123",
	})
	// 调用加密方法，发挥Token字符串
	signingString, err := claims.SignedString(key)
	if err != nil {
		return
	}
	t.Log(signingString)
}
