package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
	"Drifting/services/qiniu"
	"mime/multipart"
)

// CreateNewDriftingDrawing 创建漂流画
func CreateNewDriftingDrawing(NewDrawing model.DriftingDrawing) (error, uint) {
	err := mysql.DB.Create(&NewDrawing).Error
	if err != nil {
		return err, 0
	}
	var FindDrawing model.DriftingDrawing
	err = mysql.DB.Where(&NewDrawing).Find(&FindDrawing).Error
	if err != nil {
		return err, 0
	}
	return err, FindDrawing.ID
}

func GetDriftingDrawing(StudentID int64) ([]model.DriftingDrawing, error) {
	var drawings []model.DriftingDrawing
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&drawings).Error
	return drawings, err
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

// DrawDrawing 创作漂流画
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

// DriftingDrawingDetail 获取漂流画详情
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

// GetJoinedDriftingDrawings 获取某人加入的漂流画
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
			err = mysql.DB.Where("id = ?", v.DriftingDrawingID).First(&a).Error
			if err != nil {
				return nil, err
			}
			drawings = append(drawings, a)
		}
	}
	return drawings, nil
}

// RandomRecommendDrawing 随机推荐漂流画
func RandomRecommendDrawing() (model.DriftingDrawing, error) {
	var drawings []model.DriftingDrawing
	err := mysql.DB.Not("kind", "熟人模式").Find(&drawings).Error
	if err != nil {
		return model.DriftingDrawing{}, err
	}
	m1 := make(map[int]model.DriftingDrawing)
	for i := 0; i < len(drawings); i++ {
		m1[i] = drawings[i]
	}
	var ret model.DriftingDrawing
	for _, v := range m1 {
		ret = v
		break
	}
	for k := range m1 {
		delete(m1, k)
	}
	return ret, nil
}

// DeleteDrawing 删除漂流画
func DeleteDrawing(DLDrawing model.DriftingDrawing) error {
	err := mysql.DB.Where(&DLDrawing).Delete(&DLDrawing).Error
	return err
}
