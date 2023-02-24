package driftingpicture

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/driftingfile"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 创建漂流照片
// @Description 创建漂流照片
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewPicture body model.CreateFile true "新建漂流照片"
// @Success 200 {object} handler.Response "{"message":"创建成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_picture/create [post]
func CreateDriftingPicture(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewPicture model.DriftingPicture
	err := c.BindJSON(&NewPicture)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	NewPicture.OwnerID = StudentID
	err = driftingfile.CreateNewDriftingPicture(NewPicture)
	if err != nil {
		handler.SendBadResponse(c, "创建失败", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功", nil)
}

// @Summary 获取用户漂流照片
// @Description 获取对应用户创建的漂流照片
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingPicture "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_picture/create  [get]
func GetCreatedDriftingPictures(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	notes, err := driftingfile.GetDriftingNotes(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", notes)
}

// @Summary 参加漂流照片创作(仅参加)
// @Description 参加漂流照片创作(仅参加)
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Joining body model.JoinedDrifting true "要参加的漂流照片"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_picture/join [post]
func JoinDriftingPicture(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var Joining model.JoinedDrifting
	err := c.BindJSON(&Joining)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	Joining.StudentID = StudentID
	err = driftingfile.JoinNewDriftingPicture(Joining)
	if err != nil {
		handler.SendBadResponse(c, "创建失败", err)
		return
	}
	handler.SendGoodResponse(c, "参加成功", nil)
}

// @Summary 获取用户参加的漂流照片信息
// @Description 获取用户参加得漂流照片信息
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingNote "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_picture/join [get]
func GetJoinedDriftingPictures(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	pictures, err := driftingfile.GetJoinedDriftingPictures(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取出错", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", pictures)
}

// @Summary 创作漂流照片
// @Description 创作漂流照片
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param file formData file true "内容"
// @Param id formData string true "id"
// @Success 200 {object} handler.Response "{"message":"创建成功"}"
// @Failure 400 {object} handler.Response "{"message":"创建失败"}"
// @Router /api/v1/drifting_picture/draw [post]
func DrawDriftingPicture(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewContact model.PictureContact
	a := c.Param("file_id")
	b, err := strconv.Atoi(a)
	if err != nil {
		handler.SendBadResponse(c, "出错", err)
		return
	}
	NewContact.FileID = int64(b)
	f, err := c.FormFile("picture")
	if err != nil {
		handler.SendBadResponse(c, "出错", err)
		return
	}
	err = driftingfile.DrawPicture(StudentID, NewContact, f)
	if err != nil {
		handler.SendBadResponse(c, "创建出错", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功", nil)
}

// @Summary 获取漂流照片内容
// @Description 获取漂流本内容，需在json中提供漂流照片的ID
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param FDriftingNote body model.GetFileId true "获取的ID"
// @Success 200 {object} model.PictureInfo "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_picture/detail [get]
func GetDriftingPictureDetail(c *gin.Context) {
	var FDriftingPicture model.DriftingPicture
	info, err := driftingfile.DriftingPictureDetail(FDriftingPicture)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", info)
}

// @Summary 邀请好友进行创作
// @Description 邀请好友创作，需在json中提供好友学号，漂流本ID，及文件类型(漂流本需注明是DriftingNote)
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewInvite body model.Invite true "新建邀请"
// @Success 200 {object} handler.Response "{"message":"邀请成功"}"
// @Failure 400 {object} handler.Response "{"message":"邀请失败，你可能已邀请过该好友"}"
// @Router /api/v1/drifting_picture/invite [post]
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
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Invite "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取信息失败"}"
// @Router /api/v1/drifting_picture/invite [get]
func GetInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	invites, err := driftingfile.GetInvites(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", invites)
}

// @Summary 拒绝邀请
// @Description 拒绝创作邀请
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "拒绝邀请"
// @Success 200 {object} handler.Response "{"message":"拒绝成功"}"
// @Failure 400 {object} handler.Response "{"message":"拒绝失败"}"
// @Router /api/v1/drifting_picture/refuse [post]
func RefuseInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var TheInvite model.Invite
	err := c.BindJSON(&TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	TheInvite.FriendID = StudentID
	err = driftingfile.RefusePictureInvite(TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "拒绝操作失败", err)
		return
	}
	handler.SendGoodResponse(c, "拒绝成功", nil)
}

// @Summary 随机推荐漂流照片
// @Description 随机推荐一个漂流照片
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} model.DriftingPicture "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_picture/recommendation [get]
func RandomRecommendation(c *gin.Context) {
	TheNote, err := driftingfile.RandomRecommend()
	if err != nil {
		handler.SendBadResponse(c, "漂流照片推送失败", err)
		return
	}
	handler.SendGoodResponse(c, "推送成功", TheNote)
}

// @Summary 接受创作邀请
// @Description 接受好友创作邀请，注：该接口仅负责删除对应邀请记录，后续操作需调用参与创作接口
// @Tags driftingpicture
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "要通过的邀请"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_picture/accept [post]
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
