package router

import (
	"Drifting/handler/driftingfile/driftingnote"
	"Drifting/handler/user"
	"Drifting/handler/user/friend"
	"Drifting/router/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()
	LoginGroup := e.Group("/api/v1/login")
	{
		LoginGroup.POST("", user.Login)
	}
	//用户相关路由
	UserGroup := e.Group("/api/v1/user").Use(middleware.Auth())
	{
		UserGroup.GET("/detail", user.GetUserDetails)
		UserGroup.PUT("/update", user.UpdateUserInfo)
		UserGroup.PUT("/avatar", user.UpdateUserAvatar)
	}

	//好友相关路由
	FriendGroup := e.Group("/api/v1/friend").Use(middleware.Auth())
	{
		FriendGroup.POST("/add", friend.AddFriend)
		FriendGroup.GET("/get", friend.GetFriend)
		FriendGroup.GET("/request", friend.GetAddRequest)
		FriendGroup.POST("/pass", friend.PassAddRequest)
		FriendGroup.DELETE("", friend.DeleteFriend)
	}

	//漂流本路由
	DriftingNoteGroup := e.Group("/api/v1/driftingnote").Use(middleware.Auth())
	{
		DriftingNoteGroup.POST("/create", driftingnote.CreateDriftingNote)
	}
	return e
}
