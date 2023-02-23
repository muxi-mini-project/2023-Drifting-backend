package middleware

import (
	"Drifting/handler"
	"Drifting/pkg/auth"
	"Drifting/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaim, err := auth.ParseRequest(c)
		if err != nil {
			handler.SendError(c, errno.ErrTokenInvalid, err.Error())
			//终止函数运行
			c.Abort()
			return
		}

		// 跨越中间件取值
		c.Set("student_id", userClaim.StudentID)
		c.Set("expiresAt", userClaim.StandardClaims)

		c.Next()
	}
}
