package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
	"Drifting/services/qiniu"
	"mime/multipart"
)

// CreateDriftingNote 创建漂流本
func CreateDriftingNote(StudentID int64, NewDriftingNote model.DriftingNote, f *multipart.FileHeader) error {
	NewDriftingNote.OwnerID = StudentID
	_, url := qiniu.UploadToQiNiu(f, "note_covers/")
	NewDriftingNote.Cover = url
	err := mysql.DB.Create(&NewDriftingNote).Error
	return err
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
			err = mysql.DB.Where("drifting_note_id = ?", v.DriftingNoteID).First(&a).Error
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

// RefuseInvite 拒绝邀请
func RefuseInvite(TheInvite model.Invite) error {
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

//// RandomRecommend 随机推荐漂流本
//func RandomRecommend() (model.NoteInfo, error) {
//
//}
