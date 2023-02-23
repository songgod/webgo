package middleware

import (
	"back/common"
	"back/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 用户请求中获取token信息
		tokenString := c.GetHeader("Authorization")

		// 判断token的有效性
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足0"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		// 解析token串，返回token及claims
		t, c2, err := common.ParseToken(tokenString)
		if err != nil || !t.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			c.Abort()
			return
		}

		// 获取claim中的userid
		userId := c2.UserId

		// 查询数据库查找用户
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		// 判断用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			c.Abort()
			return
		}

		// 如果用户存在，将用户信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
