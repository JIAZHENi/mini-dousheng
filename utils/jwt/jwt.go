package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var Jwtkey = []byte("记忆")

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(userId int64, userName string) (string, error) {
	expireTime := time.Now().Add(720 * time.Hour) //过期时间 720 / 24 = 30 天
	nowTime := time.Now()                         //当前时间
	claims := MyClaims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "JIYI",            //颁发者签名
			Subject:   "userToken",       //签名主题
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString(Jwtkey)
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid && time.Now().Unix() < key.ExpiresAt {
		return key, true
	} else {
		return nil, false
	}
}
