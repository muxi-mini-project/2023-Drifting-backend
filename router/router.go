package router

import (
	"Drifting/handler/user"
	"Drifting/handler/user/friend"
	"Drifting/router/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()
	group0 := e.Group("/api/v1/login")
	{
		group0.POST("", user.Login)
	}
	//用户相关路由
	group1 := e.Group("/api/v1/user").Use(middleware.Auth())
	{
		group1.GET("/detail", user.GetUserDetails)
		group1.PUT("/update", user.UpdateUserInfo)
		group1.PUT("/avatar", user.UpdateUserAvatar)
	}

	//好友相关路由
	group2 := e.Group("/api/v1/friend").Use(middleware.Auth())
	{
		group2.POST("/friend", friend.AddFriend)
		group2.GET("/get", friend.GetFriend)
		group2.GET("/request", friend.GetAddRequest)
		group2.POST("/pass", friend.PassAddRequest)
		group2.DELETE("/delete", friend.DeleteFriend)
	}
	return e
}
