package draft

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/draft"
	"github.com/gin-gonic/gin"
)

// @Summary 创建草稿箱
// @Description 创建草稿箱
// @Tags draft
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Draft body model.CreateFile true "新建草稿箱信息"
// @Success 200 {object} handler.Response "{"message":"创建成功"}"
// @Failure 400 {object} handler.Response "{"message":"创建失败"}"
// @Router /api/v1/draft/create [post]
func CreateDraft(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewDraft model.Draft
	err := c.BindJSON(&NewDraft)
	if err != nil {
		handler.SendBadResponse(c, "获取信息出错", err)
		return
	}
	err = draft.CreateDraft(StudentID, NewDraft)
	if err != nil {
		handler.SendBadResponse(c, "创建出错", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功", nil)
}

// @Summary 加入草稿(写内容)
// @Description 加入草稿,需要在json中添加名为
// @Tags draft
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewContact body model.NoteContact true "内容"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/draft/write [post]
func WriteDraft(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewContact model.DraftContact
	err := c.BindJSON(&NewContact)
	if err != nil {
		handler.SendBadResponse(c, "获取数据出错", err)
		return
	}
	err = draft.WriteDraft(StudentID, NewContact)
	if err != nil {
		handler.SendBadResponse(c, "存储失败", err)
		return
	}
	handler.SendGoodResponse(c, "参与创作成功", nil)
}

// @Summary 获取用户草稿箱
// @Description 获取对应用户创建的草稿箱
// @Tags draft
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Draft "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/draft/create  [get]
func GetCreatedDrafts(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	notes, err := draft.GetDrafts(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", notes)
}
