package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
)

// CreateDriftingNote 创建漂流本
func CreateDriftingNote(StudentID int64, NewDriftingNote model.DriftingNote) (error, uint) {
	NewDriftingNote.OwnerID = StudentID
	//NewDriftingNote.WriterNumber = 1
	err := mysql.DB.Create(&NewDriftingNote).Error
	if err != nil {
		return err, 0
	}
	var FindNote model.DriftingNote
	err = mysql.DB.Where(&NewDriftingNote).Find(&FindNote).Error
	if err != nil {
		return err, 0
	}
	//var NewJoined model.JoinedDrifting
	//NewJoined.StudentID = StudentID
	//NewJoined.DriftingNoteID = int64(NewDriftingNote.ID)
	//err = JoinDriftingNote(NewJoined)
	//if err != nil {
	//	return err, 0
	//}
	return err, FindNote.ID
}

// WriteDriftingNote 参与创作
func WriteDriftingNote(StudentID int64, TheContact model.NoteContact) error {
	TheContact.WriterID = StudentID
	err := mysql.DB.Create(&TheContact).Error
	var NewInfo model.DriftingNote
	err = mysql.DB.Where("id = ?", TheContact.FileID).Find(&NewInfo).Error
	if err != nil {
		return err
	}
	NewInfo.WriterNumber = NewInfo.WriterNumber + 1
	err = mysql.DB.Where("id = ?", TheContact.FileID).Updates(&NewInfo).Error
	return err
}

// GetDriftingNotes 获取某人的漂流本
func GetDriftingNotes(StudentID int64) ([]model.DriftingNote, error) {
	var notes []model.DriftingNote
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&notes).Error
	return notes, err
}

// JoinDriftingNote 参加漂流本创作
func JoinDriftingNote(Joining model.JoinedDrifting) error {
	err := mysql.DB.Where(&Joining).First(&Joining).Error
	if err != nil {
		err1 := mysql.DB.Create(&Joining).Error
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
			err = mysql.DB.Where("id = ?", v.DriftingNoteID).Find(&a).Error
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

// SelectID 检验邀请ID是否匹配
func SelectID(NewInvite model.Invite, a string) error {
	switch a {
	case "漂流本":
		var NewNote model.DriftingNote
		NewNote.ID = uint(NewInvite.FileID)
		NewNote.OwnerID = NewInvite.HostID
		err := mysql.DB.Where(&NewNote).First(&NewNote).Error
		if err == nil {
			break
		} else {
			return errno.ErrMatch
		}
	case "漂流画":
		var NewDrawing model.DriftingDrawing
		NewDrawing.ID = uint(NewInvite.FileID)
		NewDrawing.OwnerID = NewInvite.HostID
		err := mysql.DB.Where(&NewDrawing).First(&NewDrawing).Error
		if err == nil {
			break
		} else {
			return errno.ErrMatch
		}
	case "漂流小说":
		var NewNovel model.DriftingNovel
		NewNovel.ID = uint(NewInvite.FileID)
		NewNovel.OwnerID = NewInvite.HostID
		err := mysql.DB.Where(&NewNovel).First(&NewNovel).Error
		if err == nil {
			break
		} else {
			return errno.ErrMatch
		}
	case "漂流相机":
		var NewPicture model.DriftingPicture
		NewPicture.ID = uint(NewInvite.FileID)
		NewPicture.OwnerID = NewInvite.HostID
		err := mysql.DB.Where(&NewPicture).First(&NewPicture).Error
		if err == nil {
			break
		} else {
			return errno.ErrMatch
		}
	}
	return nil
}

// CreateInvite 创建创作邀请
func CreateInvite(NewInvite model.Invite, a string) error {
	NewInvite.FileKind = a
	err := mysql.DB.Where(&NewInvite).First(&NewInvite).Error
	if err != nil {
		err = SelectID(NewInvite, a)
		if err == nil {
			err = mysql.DB.Create(&NewInvite).Error
		}
		return err
	}
	return errno.ErrDatabase
}

func CreateDrawingInviteInfos(info model.DriftingDrawing) model.InviteInfo {
	var ThisInfo model.InviteInfo
	ThisInfo = model.InviteInfo{
		Name:     info.Name,
		FileID:   info.ID,
		CreateAt: info.CreatedAt,
		FileKind: "漂流画",
		OwnerID:  info.OwnerID,
		Cover:    info.Cover,
		Kind:     info.Kind,
		Theme:    info.Theme,
		Number:   info.SetNumber,
	}
	return ThisInfo
}

func CreateNoteInviteInfos(info model.DriftingNote) model.InviteInfo {
	var ThisInfo model.InviteInfo
	ThisInfo = model.InviteInfo{
		Name:     info.Name,
		FileID:   info.ID,
		CreateAt: info.CreatedAt,
		FileKind: "漂流本",
		OwnerID:  info.OwnerID,
		Cover:    info.Cover,
		Kind:     info.Kind,
		Theme:    info.Theme,
		Number:   info.SetNumber,
	}
	return ThisInfo
}

func CreatePictureInviteInfos(info model.DriftingPicture) model.InviteInfo {
	var ThisInfo model.InviteInfo
	ThisInfo = model.InviteInfo{
		Name:     info.Name,
		FileID:   info.ID,
		CreateAt: info.CreatedAt,
		FileKind: "漂流相机",
		OwnerID:  info.OwnerID,
		Cover:    info.Cover,
		Kind:     info.Kind,
		Theme:    info.Theme,
		Number:   info.SetNumber,
	}
	return ThisInfo
}

func CreateNovelInviteInfos(info model.DriftingNovel) model.InviteInfo {
	var ThisInfo model.InviteInfo
	ThisInfo = model.InviteInfo{
		Name:     info.Name,
		FileID:   info.ID,
		CreateAt: info.CreatedAt,
		FileKind: "漂流小说",
		OwnerID:  info.OwnerID,
		Cover:    info.Cover,
		Kind:     info.Kind,
		Theme:    info.Theme,
		Number:   info.SetNumber,
	}
	return ThisInfo
}

// GetInvites 获取邀请信息
func GetInvites(StudentID int64, num int) ([]model.InviteInfo, error) {
	var invites []model.Invite
	var err error
	switch num {
	case 1:
		err = mysql.DB.Where("friend_id = ? AND file_kind = ?", StudentID, "漂流画").Find(&invites).Error
		break
	case 2:
		err = mysql.DB.Where("friend_id = ? AND file_kind = ?", StudentID, "漂流小说").Find(&invites).Error
		break
	case 3:
		err = mysql.DB.Where("friend_id = ? AND file_kind = ?", StudentID, "漂流本").Find(&invites).Error
		break
	case 4:
		err = mysql.DB.Where("friend_id = ? AND file_kind = ?", StudentID, "漂流相机").Find(&invites).Error
		break
	}
	if err != nil {
		return nil, err
	}
	var InviteInfos []model.InviteInfo
	for _, invite := range invites {
		if num == 1 {
			var info model.DriftingDrawing
			err = mysql.DB.Where("id = ?", invite.FileID).Find(&info).Error
			InviteInfos = append(InviteInfos, CreateDrawingInviteInfos(info))
		} else if num == 2 {
			var info model.DriftingNovel
			err = mysql.DB.Where("id = ?", invite.FileID).Find(&info).Error
			InviteInfos = append(InviteInfos, CreateNovelInviteInfos(info))
		} else if num == 3 {
			var info model.DriftingNote
			err = mysql.DB.Where("id = ?", invite.FileID).Find(&info).Error
			InviteInfos = append(InviteInfos, CreateNoteInviteInfos(info))
		} else if num == 4 {
			var info model.DriftingPicture
			err = mysql.DB.Where("id = ?", invite.FileID).Find(&info).Error
			InviteInfos = append(InviteInfos, CreatePictureInviteInfos(info))
		}
	}
	return InviteInfos, err
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
	Note.SetNumber = Note.SetNumber - 1
	err = mysql.DB.Where("id = ?", Note.ID).Updates(&Note).Error
	return err
}

// RandomRecommendNote 随机推荐漂流本
func RandomRecommendNote(StudentID int64) (model.DriftingNote, error) {
	var notes []model.DriftingNote
	err := mysql.DB.Not("kind", 1).Not("number", 1, 0).Not("owner_id", StudentID).Find(&notes).Error
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
	//if err != nil {
	//	return err
	//}
	//var NewJoined model.JoinedDrifting
	//NewJoined.StudentID = TheInvite.FriendID
	//switch TheInvite.FileKind {
	//case "漂流本":
	//	NewJoined.DriftingNoteID = TheInvite.FileID
	//	err = JoinDriftingNote(NewJoined)
	//	break
	//case "漂流画":
	//	NewJoined.DriftingDrawingID = TheInvite.FileID
	//	err = JoinNewDriftingDrawing(NewJoined)
	//	break
	//case "漂流小说":
	//	NewJoined.DriftingNovelID = TheInvite.FileID
	//	err = JoinDriftingNovel(NewJoined)
	//	break
	//case "漂流相机":
	//	NewJoined.DriftingPictureID = TheInvite.FileID
	//	err = JoinNewDriftingPicture(NewJoined)
	//	break
	//}
	return err
}

// DeleteNote 删除指定漂流本
func DeleteNote(TheNote model.DriftingNote) error {
	err := mysql.DB.Where(&TheNote).Delete(&TheNote).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Where("drifting_note_id = ?", TheNote.ID).Delete(&model.JoinedDrifting{}).Error
	return err
}
