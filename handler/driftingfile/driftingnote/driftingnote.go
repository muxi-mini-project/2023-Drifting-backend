package driftingnote

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/driftingfile"
	"github.com/gin-gonic/gin"
)

// @Summary 创建漂流本
// @Description 创建漂流本,kind必备，且只能为"熟人模式"和"生人模式"，否则将无法进行筛选及推送
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param DriftingNote body model.CreateFile true "新建漂流本信息"
// @Success 200 {object} handler.Response "{"message":"创建成功"}"
// @Failure 400 {object} handler.Response "{"message":"创建失败"}"
// @Router /api/v1/drifting_note/create [post]
func CreateDriftingNote(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewDriftingNote model.DriftingNote
	err := c.BindJSON(&NewDriftingNote)
	if err != nil {
		handler.SendBadResponse(c, "获取信息出错", err)
		return
	}
	var id uint
	err, id = driftingfile.CreateDriftingNote(StudentID, NewDriftingNote)
	if err != nil {
		handler.SendBadResponse(c, "创建出错", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功，获得漂流本id", id)
}

// @Summary 参与漂流本创作(写内容)
// @Description 参与漂流本创作,需要在json中添加名为
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewContact body model.NoteContact true "内容"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_note/write [post]
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

// @Summary 获取用户漂流本
// @Description 获取对应用户创建的漂流本
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingNote "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_note/create  [get]
func GetCreatedDriftingNotes(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	notes, err := driftingfile.GetDriftingNotes(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", notes)
}

// @Summary 参与漂流本创作(仅参与)
// @Description 参与漂流本创作
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Joined body model.JoinedDrifting true "要参加的漂流本"
// @Success 200 {object} handler.Response "{"message":"参加成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_note/join [post]
func JoinDrifting(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var Joined model.JoinedDrifting
	err := c.BindJSON(&Joined)
	if err != nil {
		handler.SendBadResponse(c, "获取安卓信息出错", err)
		return
	}
	Joined.StudentID = StudentID
	err = driftingfile.JoinDriftingNote(Joined)
	if err != nil {
		handler.SendBadResponse(c, "参与出错，请确定您是否已经参与或传入信息有误", err)
		return
	}
	handler.SendGoodResponse(c, "参加成功", nil)
}

// @Summary 获取用户参加的漂流本信息
// @Description 获取用户参加得漂流本信息
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingNote "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_note/join [get]
func GetJoinedDriftingNotes(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	notes, err := driftingfile.GetJoinedDriftingNotes(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取出错", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", notes)
}

// @Summary 获取漂流本内容
// @Description 获取漂流本内容，需在json中提供漂流本的ID
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param FDriftingNote body model.DriftingNote true "获取的ID"
// @Success 200 {object} model.NoteInfo "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_note/detail [get]
func GetDriftingNoteDetail(c *gin.Context) {
	var FDriftingNote model.DriftingNote
	c.BindJSON(&FDriftingNote)
	info, err := driftingfile.GetNoteInfo(FDriftingNote)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", info)
}

// @Summary 邀请好友进行创作
// @Description 邀请好友创作，需在json中提供好友学号，漂流本ID，及文件类型(漂流本需注明是DriftingNote)
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewInvite body model.Invite true "新建邀请"
// @Success 200 {object} handler.Response "{"message":"邀请成功"}"
// @Failure 400 {object} handler.Response "{"message":"邀请失败，你可能已邀请过该好友"}"
// @Router /api/v1/drifting_note/invite [post]
func InviteFriend(c *gin.Context) {
	var NewInvite model.Invite
	NewInvite.HostID = c.MustGet("student_id").(int64)
	err := c.BindJSON(&NewInvite)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	err = driftingfile.CreateInvite(NewInvite)
	if err != nil {
		handler.SendBadResponse(c, "邀请失败，你可能已经已经邀请过该好友", err)
		return
	}
	handler.SendGoodResponse(c, "邀请成功", nil)
}

// @Summary 获取邀请信息
// @Description 获取该用户的邀请信息
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Invite "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取信息失败"}"
// @Router /api/v1/drifting_note/invite [get]
func GetInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	invites, err := driftingfile.GetInvites(StudentID, 3)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", invites)
}

// @Summary 拒绝邀请
// @Description 拒绝创作邀请
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "拒绝邀请"
// @Success 200 {object} handler.Response "{"message":"拒绝成功"}"
// @Failure 400 {object} handler.Response "{"message":"拒绝失败"}"
// @Router /api/v1/drifting_note/refuse [post]
func RefuseInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var TheInvite model.Invite
	err := c.BindJSON(&TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	TheInvite.FriendID = StudentID
	err = driftingfile.RefuseNoteInvite(TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "拒绝操作失败", err)
		return
	}
	handler.SendGoodResponse(c, "拒绝成功", nil)
}

// @Summary 随机推荐漂流本
// @Description 随机推荐一个漂流本
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} model.DriftingNote "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_note/recommendation [get]
func RandomRecommendation(c *gin.Context) {
	TheNote, err := driftingfile.RandomRecommendNote()
	if err != nil {
		handler.SendBadResponse(c, "漂流本推送失败", err)
		return
	}
	handler.SendGoodResponse(c, "推送成功", TheNote)
}

// @Summary 接受创作邀请
// @Description 接受好友创作邀请，注：该接口仅负责删除对应邀请记录，后续操作需调用参与创作接口
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "要通过的邀请"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_note/accept [post]
func AcceptInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var TheInvite model.Invite
	err := c.BindJSON(&TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	TheInvite.FriendID = StudentID
	err = driftingfile.AcceptTheInvite(TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "出错", err)
		return
	}
	handler.SendGoodResponse(c, "通过成功", nil)
}

// @Summary 删除漂流本
// @Description 删除指定漂流本
// @Tags driftingnote
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheNote body model.DriftingNote true "要删除的漂流本"
// @Success 200 {object} handler.Response "{"message":"删除成功"}"
// @Failure 400 {object} handler.Response "{"message":"删除失败，您有可能不是该文件的主人，或者该文件不存在"}"
// @Router /api/v1/drifting_note/delete [delete]
func DeleteNote(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var DLNote model.DriftingNote
	err := c.BindJSON(&DLNote)
	DLNote.OwnerID = StudentID
	if err != nil {
		handler.SendBadResponse(c, "获取id出错", err)
		return
	}
	err = driftingfile.DeleteNote(DLNote)
	if err != nil {
		handler.SendBadResponse(c, "删除失败，您有可能不是该文件的主人，或者该文件不存在", err)
		return
	}
	handler.SendGoodResponse(c, "删除成功", err)
}
