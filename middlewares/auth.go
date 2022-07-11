package middlewares

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.获取 token
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		// 2.验证 token
		if token == "" {
			log.Println("user unLogin")
			c.Set("loginId", int64(0))
		} else {
			JWTMassage, err := jwt.CheckToken(token)
			if err != true {
				log.Println("token check false")
				model.ResponseError(c)
				c.Abort()
				return
			}
			c.Set("loginId", JWTMassage.UserId)
		}
	}
}
