package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
)

// CreateDriftingNote 创建漂流本
func CreateDriftingNote(StudentID int64, NewDriftingNote model.DriftingNote) (error, uint) {
	NewDriftingNote.OwnerID = StudentID
	err := mysql.DB.Create(&NewDriftingNote).Error
	if err != nil {
		return err, 0
	}
	var FindNote model.DriftingNote
	err = mysql.DB.Where(&NewDriftingNote).Find(&FindNote).Error
	if err != nil {
		return err, 0
	}
	return err, FindNote.ID
}

// WriteDriftingNote 参与创作
func WriteDriftingNote(StudentID int64, TheContact model.NoteContact) error {
	TheContact.WriterID = StudentID
	err := mysql.DB.Create(&TheContact).Error
	return err
}

// GetDriftingNotes 获取某人的漂流本
func GetDriftingNotes(StudentID int64) ([]model.DriftingNote, error) {
	var notes []model.DriftingNote
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&notes).Error
	return notes, err
}

// JoinDriftingNote 参加漂流本创作
func JoinDriftingNote(NewJoin model.JoinedDrifting) error {
	err := mysql.DB.Where(&NewJoin).First(&NewJoin).Error
	if err != nil {
		err1 := mysql.DB.Create(&NewJoin).Error
		return err1
	}
	return errno.ErrDatabase
}

// GetJoinedDriftingNotes 获取某人加入的漂流本
func GetJoinedDriftingNotes(StudentID int64) ([]model.DriftingNote, error) {
	var notes []model.DriftingNote
	var Joined []model.JoinedDrifting
	err := mysql.DB.Where("student_id = ?", StudentID).Find(&Joined).Error
	if err != nil {
		return nil, err
	}
	for _, v := range Joined {
		if v.DriftingNoteID != 0 {
			var a model.DriftingNote
			err = mysql.DB.Where("id = ?", v.DriftingNoteID).First(&a).Error
			if err != nil {
				return nil, err
			}
			notes = append(notes, a)
		}
	}
	return notes, nil
}

// GetNoteInfo 获取漂流本内容
func GetNoteInfo(FD model.DriftingNote) (model.NoteInfo, error) {
	var info model.NoteInfo
	err := mysql.DB.Where(&FD).First(&FD).Error
	if err != nil {
		return model.NoteInfo{}, err
	}
	info.Name = FD.Name
	info.OwnerID = FD.OwnerID
	err = mysql.DB.Where("file_id = ?", FD.ID).Find(&info.Contacts).Error
	if err != nil {
		return model.NoteInfo{}, err
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

// RefuseNoteInvite 拒绝漂流本邀请
func RefuseNoteInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	if err != nil {
		return err
	}
	var Note model.DriftingNote
	err = mysql.DB.Where("id = ?", TheInvite.FileID).First(&Note).Error
	if err != nil {
		return err
	}
	Note.Number = Note.Number - 1
	err = mysql.DB.Where("id = ?", Note.ID).Updates(&Note).Error
	return err
}

// RandomRecommendNote 随机推荐漂流本
func RandomRecommendNote() (model.DriftingNote, error) {
	var notes []model.DriftingNote
	err := mysql.DB.Not("kind", "熟人模式").Find(&notes).Error
	if err != nil {
		return model.DriftingNote{}, err
	}
	m1 := make(map[int]model.DriftingNote)
	for i := 0; i < len(notes); i++ {
		m1[i] = notes[i]
	}
	var ret model.DriftingNote
	for _, v := range m1 {
		ret = v
		break
	}
	for k := range m1 {
		delete(m1, k)
	}
	return ret, nil
}

// AcceptTheInvite 接受邀请
func AcceptTheInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	return err
}
