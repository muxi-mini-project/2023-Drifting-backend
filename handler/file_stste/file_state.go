package state

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/file_state"
	"Drifting/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary 对漂流文件上锁
// @Description 需提供文件ID，及文件类型(漂流本/漂流照片/漂流小说/漂流画)，需携带token
// @Tags Lock
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Request body model.State true "上锁请求"
// @Success 200 {object} handler.Response "{"message":"上锁成功"}"
// @Failure 400 {object} handler.Response "{"message":"上锁失败"}"
// @Router /api/v1/lock/lock_on [post]
func LockOnDrifting(c *gin.Context) {
	var LockRequest model.State
	err := c.BindJSON(&LockRequest)
	LockRequest.WriterID = c.MustGet("student_id").(int64)
	if err != nil {
		handler.SendBadResponse(c, "获取信息出错", err)
		return
	}
	err, Final := file_state.LockOn(LockRequest)
	if err != nil {
		if err == errno.ErrState {
			handler.SendBadResponse(c, "该文件正有人在创作", Final)
			return
		}
		handler.SendBadResponse(c, "上锁失败", err)
		return
	}
	handler.SendGoodResponse(c, "上锁成功，开始创作", err)
}

// @Summary 解锁漂流文件
// @Description 解锁漂流文件，需提供文件ID及文件类型(漂流本/漂流照片/漂流小说/漂流画)
// @Tags Lock
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Request body model.State true "解锁请求体"
// @Success 200 {object} handler.Response "{"message":"解锁成功"}"
// @Failure 400 {object} handler.Response "{"message":"解锁失败"}"
// @Router /api/v1/lock/lock_off [delete]
func UnlockDrifting(c *gin.Context) {
	var UnlockRequest model.State
	err := c.BindJSON(&UnlockRequest)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	UnlockRequest.WriterID = c.MustGet("student_id").(int64)
	err = file_state.UnLock(UnlockRequest)
	if err != nil {
		if err == errno.ErrUnLock {
			handler.SendBadResponse(c, "你不是该文件的上锁者", err)
			return
		}
		handler.SendBadResponse(c, "解锁失败", err)
		return
	}
	handler.SendGoodResponse(c, "解锁成功", err)
}

// @Summary 或取当前上锁人
// @Description 获取当前上锁人信息，需提供文件ID及文件类型(漂流本/漂流照片/漂流小说/漂流画)
// @Tags Lock
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Request body model.State true "获取信息请求"
// @Success 200 {object} handler.Response "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/lock/get_lock [post]
func GetLock(c *gin.Context) {
	var ViewRequest model.State
	err := c.BindJSON(&ViewRequest)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	err, ReturnInfo := file_state.GetLock(ViewRequest)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取信息成功", ReturnInfo)
}
