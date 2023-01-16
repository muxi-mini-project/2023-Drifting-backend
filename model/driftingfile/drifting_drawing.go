package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
	"Drifting/services/qiniu"
	"mime/multipart"
)

// CreateNewDriftingDrawing 创建漂流画
func CreateNewDriftingDrawing(NewDrawing model.DriftingDrawing) error {
	err := mysql.DB.Create(&NewDrawing).Error
	return err
}

// JoinNewDriftingDrawing 参加漂流画
func JoinNewDriftingDrawing(Joining model.JoinedDrifting) error {
	err := mysql.DB.Where(&Joining).First(&Joining).Error
	if err != nil {
		err = mysql.DB.Create(&Joining).Error
		return err
	}
	return errno.ErrDatabase
}

// DrawiDrawing 创作漂流画
func DrawDrawing(StudentID int64, NewDrawingContact model.DrawingContact, f *multipart.FileHeader) error {
	NewDrawingContact.WriterID = StudentID
	_, url := qiniu.UploadToQiNiu(f, "drifting_drawing/")
	NewDrawingContact.Picture = url
	err := mysql.DB.Create(&NewDrawingContact).Error
	return err
}

// RefuseDrawingInvite 拒绝漂流画邀请
func RefuseDrawingInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	if err != nil {
		return err
	}
	var Note model.DriftingDrawing
	err = mysql.DB.Where("id = ?", TheInvite.FileID).First(&Note).Error
	if err != nil {
		return err
	}
	Note.Number = Note.Number - 1
	err = mysql.DB.Where("id = ?", Note.ID).Updates(&Note).Error
	return err
}

func DriftingDrawingDetail(FD model.DriftingDrawing) (model.DrawingInfo, error) {
	var info model.DrawingInfo
	err := mysql.DB.Where(&FD).First(&FD).Error
	if err != nil {
		return model.DrawingInfo{}, err
	}
	info.Name = FD.Name
	info.OwnerID = FD.OwnerID
	err = mysql.DB.Where("file_id = ?", FD.ID).Find(&info.Contacts).Error
	if err != nil {
		return model.DrawingInfo{}, err
	}
	return info, nil
}

// GetJoinedDriftingDrawings 获取某人加入的漂流本
func GetJoinedDriftingDrawings(StudentID int64) ([]model.DriftingDrawing, error) {
	var drawings []model.DriftingDrawing
	var Joined []model.JoinedDrifting
	err := mysql.DB.Where("student_id = ?", StudentID).Find(&Joined).Error
	if err != nil {
		return nil, err
	}
	for _, v := range Joined {
		if v.DriftingNoteID != 0 {
			var a model.DriftingDrawing
			err = mysql.DB.Where("id = ?", v.DriftingNoteID).First(&a).Error
			if err != nil {
				return nil, err
			}
			drawings = append(drawings, a)
		}
	}
	return drawings, nil
}
