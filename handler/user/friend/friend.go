package friend

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/user"
	"github.com/gin-gonic/gin"
)

// @Summary 建立好友申请
// @Description 建立好友申请
// @Tags friend
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Adding body model.AddingFriend true "好友申请"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/friend/add [post]
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
	handler.SendGoodResponse(c, "申请已发出", nil)
}

// @Summary 获取好友列表信息
// @Description 获取好友列表信息
// @Tags friend
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.UserInfo "{"msg":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取好友信息出错"}"
// @Router /api/v1/friend/get [get]
func GetFriend(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	FriendsInfo, err := user.GetFriend(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取好友信息出错！", err.Error())
		return
	}
	handler.SendGoodResponse(c, "获取成功", FriendsInfo)
}

// @Summary 获取好友申请
// @Description 获取好友申请
// @Tags friend
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.UserInfo "{"msg":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/friend/request [get]
func GetAddRequest(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	FriendsInfo, err := user.GetRequest(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", nil)
		return
	}
	handler.SendGoodResponse(c, "获取成功", FriendsInfo)
}

// @Summary 通过好友申请
// @Description 通过好友申请，需将添加者的学号放在json中，对应键名为"adderID"
// @Tags friend
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param UserAndFriends body model.AddingFriend true "通过的好友学号"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/friend/pass [post]
func PassAddRequest(c *gin.Context) {
	var Adding model.AddingFriend
	Adding.TargetID = c.MustGet("student_id").(int64)
	err := c.BindJSON(&Adding)
	if err != nil {
		handler.SendBadResponse(c, "获取数据出错", err)
		return
	}
	err = user.PassRequest(Adding.AdderID, Adding.TargetID)
	if err != nil {
		handler.SendBadResponse(c, "出错", err)
		return
	}
	handler.SendGoodResponse(c, "您已通过了好友申请", nil)
}

// @Summary 删除好友
// @Description 删除对应好友，需在json中提供对应好友学号，对应键名为"friendID"
// @Tags friend
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param UserAndFriends body model.UserAndFriends true "要删除的好友"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/friend/delete [delete]
func DeleteFriend(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var UserAndFriend model.UserAndFriends
	err := c.BindJSON(&UserAndFriend)
	if err != nil {
		handler.SendBadResponse(c, "获取数据失败", err)
	}
	err1, err2 := user.Delete(StudentID, UserAndFriend.FriendId)
	if err1 != nil {
		handler.SendBadResponse(c, "删除出错", err1)
		return
	}
	if err2 != nil {
		handler.SendBadResponse(c, "删除出错", err2)
		return
	}
	handler.SendGoodResponse(c, "删除成功", nil)
}
