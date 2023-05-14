package model

import (
	"github.com/golang-jwt/jwt/v4"
	"jyksServer/common"
)

// 获取jwt
func GetJwt(user User) string {

	// 创建Token结构体
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"id":       user.Id,
	})
	// 调用加密方法，发挥Token字符串
	signingString, err := claims.SignedString(common.JwtSecretByte)
	if err != nil {
		return ""
	}
	return signingString
}
