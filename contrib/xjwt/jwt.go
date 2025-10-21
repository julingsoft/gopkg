package xjwt

import (
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id   uint64 `json:"id"`   // 用户标识
	Type int    `json:"type"` // 用户类型
}

type JWTClaims struct {
	User *User
	jwt.RegisteredClaims
}

func CreateToken(userId uint64, userType int, secretKey string) (string, error) {
	claims := &JWTClaims{
		User: &User{
			Id:   userId,
			Type: userType,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string, secretKey string) (*User, error) {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims := JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "Invalid token")
	}

	return claims.User, nil
}
