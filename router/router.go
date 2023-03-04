package router

import (
	"Drifting/handler/draft"
	"Drifting/handler/driftingfile/driftingdrawing"
	"Drifting/handler/driftingfile/driftingnote"
	"Drifting/handler/driftingfile/driftingpicture"
	driftingnovel "Drifting/handler/driftingfile/driftnovel"
	"Drifting/handler/user"
	"Drifting/handler/user/friend"
	"Drifting/router/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()
	LoginGroup := e.Group("/api/v1/login")
	{
		LoginGroup.POST("", user.Login) //一站式登录
	}
	//用户相关路由
	UserGroup := e.Group("/api/v1/user").Use(middleware.Auth())
	{
		UserGroup.GET("/detail", user.GetUserDetails)   //获取用户信息
		UserGroup.PUT("/update", user.UpdateUserInfo)   //更新用户信息
		UserGroup.PUT("/avatar", user.UpdateUserAvatar) //更新用户头像
	}

	//好友相关路由
	FriendGroup := e.Group("/api/v1/friend").Use(middleware.Auth())
	{
		FriendGroup.POST("/add", friend.AddFriend)        //添加好友
		FriendGroup.GET("/get", friend.GetFriend)         //获取好友列表
		FriendGroup.GET("/request", friend.GetAddRequest) //获取好友请求
		FriendGroup.POST("/pass", friend.PassAddRequest)  //通过好友请求
		FriendGroup.DELETE("", friend.DeleteFriend)       //删除好友
	}

	//漂流本路由

	DriftingNoteGroup := e.Group("/api/v1/drifting_note").Use(middleware.Auth())
	{
		DriftingNoteGroup.POST("/create", driftingnote.CreateDriftingNote)          //创建漂流本*
		DriftingNoteGroup.POST("/write", driftingnote.WriteDriftingNote)            //参与漂流本创作(写内容)*
		DriftingNoteGroup.GET("/create", driftingnote.GetCreatedDriftingNotes)      //获取用户创建的漂流本*
		DriftingNoteGroup.POST("/join", driftingnote.JoinDrifting)                  //参加漂流本创作(加入)*
		DriftingNoteGroup.GET("/join", driftingnote.GetJoinedDriftingNotes)         //获取参与的漂流本*
		DriftingNoteGroup.GET("/detail", driftingnote.GetDriftingNoteDetail)        //获取漂流本详情*
		DriftingNoteGroup.POST("/invite", driftingnote.InviteFriend)                //邀请好友创作*
		DriftingNoteGroup.GET("/invite", driftingnote.GetInvite)                    //获取邀请信息*
		DriftingNoteGroup.POST("/refuse", driftingnote.RefuseInvite)                //拒绝创作邀请*
		DriftingNoteGroup.POST("/accept", driftingnote.AcceptInvite)                //接受创作邀请*
		DriftingNoteGroup.GET("/recommendation", driftingnote.RandomRecommendation) //随机推送*
		DriftingNoteGroup.DELETE("/delete", driftingnote.DeleteNote)                //删除漂流本
	}

	//漂流画路由
	DriftingDrawingGroup := e.Group("/api/v1/drifting_drawing").Use(middleware.Auth())
	{
		DriftingDrawingGroup.POST("/create", driftingdrawing.CreateDriftingDrawing)       //创建漂流画
		DriftingDrawingGroup.POST("/draw", driftingdrawing.DrawDriftingDrawing)           //创作漂流画
		DriftingDrawingGroup.POST("/join", driftingdrawing.JoinDriftingDrawing)           //参加漂流画创作(仅参加)
		DriftingDrawingGroup.GET("/create", driftingdrawing.GetCreatedDriftingDrawings)   //获取用户创建的漂流画
		DriftingDrawingGroup.GET("/join", driftingdrawing.GetJoinedDriftingDrawings)      //获取用户参与的漂流画
		DriftingDrawingGroup.GET("/detail", driftingdrawing.GetDriftingDrawingDetail)     //获取漂流画信息
		DriftingDrawingGroup.POST("/invite", driftingdrawing.InviteFriend)                //邀请好友创作
		DriftingDrawingGroup.GET("/invite", driftingdrawing.GetInvite)                    //获取邀请信息
		DriftingDrawingGroup.POST("/refuse", driftingdrawing.RefuseInvite)                //拒绝创作邀请
		DriftingDrawingGroup.POST("/accept", driftingdrawing.AcceptInvite)                //接受创作邀请
		DriftingDrawingGroup.GET("/recommendation", driftingdrawing.RandomRecommendation) //随机推送
		DriftingDrawingGroup.DELETE("/delete", driftingdrawing.DeleteDrawing)             //删除漂流画
	}

	//漂流相机路由
	DriftingPictureGroup := e.Group("/api/v1/drifting_picture").Use(middleware.Auth())
	{
		DriftingPictureGroup.POST("/create", driftingpicture.CreateDriftingPicture)
		DriftingPictureGroup.GET("/create", driftingpicture.GetCreatedDriftingPictures)
		DriftingPictureGroup.POST("/join", driftingpicture.JoinDriftingPicture)
		DriftingPictureGroup.GET("/join", driftingpicture.GetJoinedDriftingPictures)
		DriftingPictureGroup.POST("/draw", driftingpicture.DrawDriftingPicture)
		DriftingPictureGroup.GET("/detail", driftingpicture.GetDriftingPictureDetail)
		DriftingPictureGroup.POST("/invite", driftingpicture.InviteFriend)
		DriftingPictureGroup.GET("/invite", driftingpicture.GetInvite)
		DriftingPictureGroup.POST("/refuse", driftingpicture.RefuseInvite)
		DriftingPictureGroup.GET("/recommendation", driftingpicture.RandomRecommendation)
		DriftingPictureGroup.POST("/accept", driftingpicture.AcceptInvite)
		DriftingPictureGroup.DELETE("/delete", driftingpicture.DeletePicture)
	}

	// 漂流小说路由
	DriftingNovelGroup := e.Group("/api/v1/drifting_novel").Use(middleware.Auth())
	{
		DriftingNovelGroup.POST("/create", driftingnovel.CreateDriftingNovel)
		DriftingNovelGroup.POST("/write", driftingnovel.WriteDriftingNovel)
		DriftingNovelGroup.GET("/create", driftingnovel.GetCreatedDriftingNovels)
		DriftingNovelGroup.POST("/join", driftingnovel.JoinDrifting)
		DriftingNovelGroup.GET("/join", driftingnovel.GetJoinedDriftingNovels)
		DriftingNovelGroup.GET("/detail", driftingnovel.GetDriftingNovelDetail)
		DriftingNovelGroup.POST("/invite", driftingnovel.InviteFriend)
		DriftingNovelGroup.GET("/invite", driftingnovel.GetInvite)
		DriftingNovelGroup.POST("/refuse", driftingnovel.RefuseInvite)
		DriftingNovelGroup.GET("/recommendation", driftingnovel.RandomRecommendation)
		DriftingNovelGroup.POST("/accept", driftingnovel.AcceptInvite)
		DriftingNovelGroup.DELETE("/delete", driftingnovel.DeletePicture)
	}

	// 草稿箱路由
	DraftGroup := e.Group("/api/v1/draft").Use(middleware.Auth())
	{
		DraftGroup.POST("/create", draft.CreateDraft)
		DraftGroup.POST("/write", draft.WriteDraft)
		DraftGroup.GET("/create", draft.GetCreatedDrafts)
		DraftGroup.DELETE("/delete", draft.DeleteDraft)
	}

	//e.POST("api/v1/test", func(c *gin.Context) {
	//	f, err := c.FormFile("picture")
	//	if err != nil {
	//		handler.SendBadResponse(c, "出错", err)
	//		return
	//	}
	//	_, str := qiniu.UploadToQiNiu(f, "covers/")
	//	handler.SendGoodResponse(c, "获取成功", str)
	//})
	return e
}
