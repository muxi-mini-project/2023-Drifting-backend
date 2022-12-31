package router

import (
	"Drifting/handler/user"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()

	group1 := e.Group("/api/v1/user")
	{
		group1.GET("/detail", user.GetUserDetails)
		group1.PUT("/update", user.UpdateUserInfo)
		group1.PUT("/avatar", user.UpdateUserAvatar)
	}

	return e
}
