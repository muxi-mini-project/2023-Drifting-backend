package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
)

// CreateDriftingNovel 创建漂流小说
func CreateDriftingNovel(StudentID int64, NewDriftingNovel model.DriftingNovel) error {
	NewDriftingNovel.OwnerID = StudentID
	err := mysql.DB.Create(&NewDriftingNovel).Error
	return err
}

// WriteDriftingNovel 参与创作
func WriteDriftingNovel(StudentID int64, TheContact model.NovelContact) error {
	TheContact.WriterID = StudentID
	err := mysql.DB.Create(&TheContact).Error
	return err
}

// GetDriftingNovels 获取某人的漂流小说
func GetDriftingNovels(StudentID int64) ([]model.DriftingNovel, error) {
	var novels []model.DriftingNovel
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&novels).Error
	return novels, err
}

// JoinDriftingNovel 参加漂流小说创作
func JoinDriftingNovel(NewJoin model.JoinedDrifting) error {
	err := mysql.DB.Where(&NewJoin).First(&NewJoin).Error
	if err != nil {
		err1 := mysql.DB.Create(&NewJoin).Error
		return err1
	}
	return errno.ErrDatabase
}

// GetJoinedDriftingNovels 获取某人加入的漂流小说
func GetJoinedDriftingNovels(StudentID int64) ([]model.DriftingNovel, error) {
	var novels []model.DriftingNovel
	var Joined []model.JoinedDrifting
	err := mysql.DB.Where("student_id = ?", StudentID).Find(&Joined).Error
	if err != nil {
		return nil, err
	}
	for _, v := range Joined {
		if v.DriftingNovelID != 0 {
			var a model.DriftingNovel
			err = mysql.DB.Where("id = ?", v.DriftingNovelID).First(&a).Error
			if err != nil {
				return nil, err
			}
			novels = append(novels, a)
		}
	}
	return novels, nil
}

// GetNovelInfo 获取漂流小说内容
func GetNovelInfo(FD model.DriftingNovel) (model.NovelInfo, error) {
	var info model.NovelInfo
	err := mysql.DB.Where(&FD).First(&FD).Error
	if err != nil {
		return model.NovelInfo{}, err
	}
	info.Name = FD.Name
	info.OwnerID = FD.OwnerID
	err = mysql.DB.Where("file_id = ?", FD.ID).Find(&info.Contacts).Error
	if err != nil {
		return model.NovelInfo{}, err
	}
	return info, nil
}

// CreateInvite 创建创作邀请
func CreateInvite(NewInvite model.Invite) error {
	err := mysql.DB.Where(&NewInvite).First(&NewInvite).Error
	if err != nil {
		err = mysql.DB.Create(&NewInvite).Error
		return err
	}
	return errno.ErrDatabase
}

// GetInvites 获取邀请信息
func GetInvites(StudentID int64) ([]model.Invite, error) {
	var invites []model.Invite
	err := mysql.DB.Where("friend_id = ?", StudentID).Find(&invites).Error
	return invites, err
}

// RefuseInvite 拒绝漂流小说邀请
func RefuseNovelInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	if err != nil {
		return err
	}
	var Novel model.DriftingNovel
	err = mysql.DB.Where("id = ?", TheInvite.FileID).First(&Novel).Error
	if err != nil {
		return err
	}
	Novel.Number = Novel.Number - 1
	err = mysql.DB.Where("id = ?", Novel.ID).Updates(&Novel).Error
	return err
}

// RandomRecommend 随机推荐漂流小说
func RandomRecommend() (model.DriftingNovel, error) {
	var novels []model.DriftingNovel
	err := mysql.DB.Not("kind", "熟人模式").Find(&novels).Error
	if err != nil {
		return model.DriftingNovel{}, err
	}
	m1 := make(map[int]model.DriftingNovel)
	for i := 0; i < len(novels); i++ {
		m1[i] = novels[i]
	}
	var ret model.DriftingNovel
	for _, v := range m1 {
		ret = v
		break
	}
	for k, _ := range m1 {
		delete(m1, k)
	}
	return ret, nil
}

// AcceptTheInvite 接受邀请
func AcceptTheInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	return err
}
