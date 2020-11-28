package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"server/extend/conf"
	"time"
)

type CustomClaims struct {
	UserID uint
	Username string
	jwt.StandardClaims
}

//生成token
func GenerateToken(id uint, username string) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := CustomClaims{
		UserID: id,
		Username: username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lomogo",
		},
	}
	//SigningMethodHS256 注意这里的加密算法，不同的加密算法对应的签名key类型不一致
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenStr , err := token.SignedString([]byte(conf.JwtConf.Key))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

//验证token是否有效
func ParseToken(tokenStr string) (*CustomClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtConf.Key), nil
	})
	if err != nil {
		log.Println("解析token失败", err)
		return nil, false
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
