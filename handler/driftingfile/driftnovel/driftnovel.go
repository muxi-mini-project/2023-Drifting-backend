package driftingnovel

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/driftingfile"

	"github.com/gin-gonic/gin"
)

// @Summary 创建漂流小说
// @Description 创建漂流小说,kind必备，且只能为"熟人模式"和"生人模式"，否则将无法进行筛选及推送
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param DriftingNovel body model.CreateFile true "新建漂流小说信息"
// @Success 200 {object} handler.Response "{"message":"创建成功"}"
// @Failure 400 {object} handler.Response "{"message":"创建失败"}"
// @Router /api/v1/drifting_novel/create [post]
func CreateDriftingNovel(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewDriftingNovel model.DriftingNovel
	err := c.BindJSON(&NewDriftingNovel)
	if err != nil {
		handler.SendBadResponse(c, "获取信息出错", err)
		return
	}
	var id uint
	err, id = driftingfile.CreateDriftingNovel(StudentID, NewDriftingNovel)
	if err != nil {
		handler.SendBadResponse(c, "创建出错", err)
		return
	}
	handler.SendGoodResponse(c, "创建成功，获得漂流小说id", id)
}

// @Summary 参与漂流小说创作(写内容)
// @Description 参与漂流小说创作,需要在json中添加名为
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewContact body model.NovelContact true "写的内容"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_novel/write [post]
func WriteDriftingNovel(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var NewContact model.NovelContact
	err := c.BindJSON(&NewContact)
	if err != nil {
		handler.SendBadResponse(c, "获取数据出错", err)
		return
	}
	err = driftingfile.WriteDriftingNovel(StudentID, NewContact)
	if err != nil {
		handler.SendBadResponse(c, "存储失败", err)
		return
	}
	handler.SendGoodResponse(c, "参与创作成功", nil)
}

// @Summary 获取用户漂流小说
// @Description 获取对应用户创建的漂流小说
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingNovel "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_novel/create  [get]
func GetCreatedDriftingNovels(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	novels, err := driftingfile.GetDriftingNovels(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", novels)
}

// @Summary 参加漂流小说创作
// @Description 参加漂流小说创作
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param Joined body model.JoinedDrifting true "要参加的漂流小说"
// @Success 200 {object} handler.Response "{"message":"参加成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_novel/join [post]
func JoinDrifting(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var Joined model.JoinedDrifting
	err := c.BindJSON(&Joined)
	if err != nil {
		handler.SendBadResponse(c, "获取安卓信息出错", err)
		return
	}
	Joined.StudentID = StudentID
	err = driftingfile.JoinDriftingNovel(Joined)
	if err != nil {
		handler.SendBadResponse(c, "参加出错，请确定您是否已经参与或传入信息有误", err)
		return
	}
	handler.SendGoodResponse(c, "参加成功", nil)
}

// @Summary 获取用户参加的漂流小说信息
// @Description 获取用户参加得漂流小说信息
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.DriftingNovel "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_novel/join [get]
func GetJoinedDriftingNovels(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	novels, err := driftingfile.GetJoinedDriftingNovels(StudentID)
	if err != nil {
		handler.SendBadResponse(c, "获取出错", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", novels)
}

// @Summary 获取漂流小说内容
// @Description 获取漂流小说内容，需在json中提供漂流小说的ID
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param FDriftingNovel body model.DriftingNovel true "获取的ID"
// @Success 200 {object} model.NovelInfo "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_novel/detail [get]
func GetDriftingNovelDetail(c *gin.Context) {
	var FDriftingNovel model.DriftingNovel
	info, err := driftingfile.GetNovelInfo(FDriftingNovel)
	if err != nil {
		handler.SendBadResponse(c, "获取失败", err)
		return
	}
	handler.SendGoodResponse(c, "获取成功", info)
}

// @Summary 邀请好友进行创作
// @Description 邀请好友创作，需在json中提供好友学号，漂流小说ID，及文件类型(漂流小说需注明是DriftingNovel)
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param NewInvite body model.Invite true "新建邀请"
// @Success 200 {object} handler.Response "{"message":"邀请成功"}"
// @Failure 400 {object} handler.Response "{"message":"邀请失败，你可能已邀请过该好友"}"
// @Router /api/v1/drifting_novel/invite [post]
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
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Invite "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取信息失败"}"
// @Router /api/v1/drifting_novel/invite [get]
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
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "拒绝邀请"
// @Success 200 {object} handler.Response "{"message":"拒绝成功"}"
// @Failure 400 {object} handler.Response "{"message":"拒绝失败"}"
// @Router /api/v1/drifting_novel/refuse [post]
func RefuseInvite(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var TheInvite model.Invite
	err := c.BindJSON(&TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "获取信息失败", err)
		return
	}
	TheInvite.FriendID = StudentID
	err = driftingfile.RefuseNovelInvite(TheInvite)
	if err != nil {
		handler.SendBadResponse(c, "拒绝操作失败", err)
		return
	}
	handler.SendGoodResponse(c, "拒绝成功", nil)
}

// @Summary 随机推荐漂流小说
// @Description 随机推荐一个漂流小说
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} model.DriftingNovel "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"获取失败"}"
// @Router /api/v1/drifting_novel/recommendation [get]
func RandomRecommendation(c *gin.Context) {
	TheNovel, err := driftingfile.RandomRecommendNovel()
	if err != nil {
		handler.SendBadResponse(c, "漂流小说推送失败", err)
		return
	}
	handler.SendGoodResponse(c, "推送成功", TheNovel)
}

// @Summary 接受创作邀请
// @Description 接受好友创作邀请，注：该接口仅负责删除对应邀请记录，后续操作需调用参与创作接口
// @Tags driftingnovel
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param TheInvite body model.Invite true "要通过的邀请"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/drifting_novel/accept [post]
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
