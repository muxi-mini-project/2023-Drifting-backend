package friend

import (
	"Drifting/controller/user"
	"Drifting/handler"
	"Drifting/model"
	"github.com/gin-gonic/gin"
)

// @Summary 建立好友申请
// @Description 建立好友申请
// @Tags friend
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Adding body model.AddingFriend true "好友申请"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failure"
// @Router api/v1/friend/add [post]
func AddFriend(c *gin.Context) {
	var Adding model.AddingFriend
	Adding.AdderID = c.MustGet("student_id").(int64)
	err := c.BindJSON(&Adding)
	if err != nil {
		return
	}
	err1 := user.SearchAddFriend(Adding.AdderID, Adding.TargetID)
	err2 := user.SearchFriends(Adding.AdderID, Adding.TargetID)
	if err1 == nil {
		handler.SendGoodResponse(c, "您已添加过此用户,请耐心等待对方通过", nil)
		return
	}
	if err2 == nil {
		handler.SendGoodResponse(c, "您已成为该用户的好友，请不要重复添加", nil)
		return
	}
	err = user.AddFriend(Adding.AdderID, Adding.TargetID)
	if err != nil {
		handler.SendBadResponse(c, "添加出错", nil)
	}
}

// @Summary 获取好友列表信息
// @Description 获取好友列表信息
// @Tags friend
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.UserInfo "{"msg":"获取成功"}"
// @Failure 400 {string} string "Failure"
// @Router api/v1/friend/get [get]
func GetFriend(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	FriendsInfo, err := user.GetFriend(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取好友信息出错！", err.Error())
		return
	}
	handler.SendGoodResponse(c, "获取成功", FriendsInfo)
}
