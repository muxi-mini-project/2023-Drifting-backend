package driftingnote

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/driftingfile"
	"github.com/gin-gonic/gin"
)

// @Summary 创建漂流本
// @Description 创建漂流本
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param DriftingNote body model.DriftingNote true "新建漂流本信息"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failure"
// @Router api/v1/driftingnote/create [post]
func CreateDriftingNote(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewDriftingNote model.DriftingNote
	err := c.BindJSON(&NewDriftingNote)
	if err != nil {
		handler.SendBadResponse(c, "获取信息出错", err)
		return
	}
	err = driftingfile.CreateDriftingNote(StudentID, NewDriftingNote)
	if err != nil {
		handler.SendBadResponse(c, "创建出错", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功", nil)
}

func WriteDriftingNote(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewContact model.NoteContact
	err := c.BindJSON(&NewContact)
	if err != nil {
		handler.SendBadResponse(c, "获取数据出错", err)
		return
	}
	err = driftingfile.WriteDriftingNote(StudentID, NewContact)
	if err != nil {
		handler.SendBadResponse(c, "存储失败", err)
		return
	}
	handler.SendGoodResponse(c, "参与创作成功", nil)
}
