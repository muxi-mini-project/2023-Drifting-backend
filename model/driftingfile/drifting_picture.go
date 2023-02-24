package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
	"Drifting/services/qiniu"
	"mime/multipart"
)

// CreateNewDriftingPicture 创建漂流照片
func CreateNewDriftingPicture(NewPicture model.DriftingPicture) error {
	err := mysql.DB.Create(&NewPicture).Error
	return err
}

// JoinNewDriftingPicture 参与漂流照片
func JoinNewDriftingPicture(Joining model.JoinedDrifting) error {
	err := mysql.DB.Where(&Joining).First(&Joining).Error
	if err != nil {
		err = mysql.DB.Create(&Joining).Error
		return err
	}
	return errno.ErrDatabase
}

// DrawiPicture 创作漂流照片
func DrawPicture(StudentID int64, NewPictureContact model.PictureContact, f *multipart.FileHeader) error {
	NewPictureContact.WriterID = StudentID
	_, url := qiniu.UploadToQiNiu(f, "drifting_picture/")
	NewPictureContact.Picture = url
	err := mysql.DB.Create(&NewPictureContact).Error
	return err
}

// RefusePictureInvite 拒绝漂流照片邀请
func RefusePictureInvite(TheInvite model.Invite) error {
	err := mysql.DB.Where(&TheInvite).Delete(&TheInvite).Error
	if err != nil {
		return err
	}
	var Note model.DriftingPicture
	err = mysql.DB.Where("id = ?", TheInvite.FileID).First(&Note).Error
	if err != nil {
		return err
	}
	Note.Number = Note.Number - 1
	err = mysql.DB.Where("id = ?", Note.ID).Updates(&Note).Error
	return err
}

func DriftingPictureDetail(FD model.DriftingPicture) (model.PictureInfo, error) {
	var info model.PictureInfo
	err := mysql.DB.Where(&FD).First(&FD).Error
	if err != nil {
		return model.PictureInfo{}, err
	}
	info.Name = FD.Name
	info.OwnerID = FD.OwnerID
	err = mysql.DB.Where("file_id = ?", FD.ID).Find(&info.Contacts).Error
	if err != nil {
		return model.PictureInfo{}, err
	}
	return info, nil
}

// GetJoinedDriftingPictures 获取某人加入的漂流照片
func GetJoinedDriftingPictures(StudentID int64) ([]model.DriftingPicture, error) {
	var pictures []model.DriftingPicture
	var Joined []model.JoinedDrifting
	err := mysql.DB.Where("student_id = ?", StudentID).Find(&Joined).Error
	if err != nil {
		return nil, err
	}
	for _, v := range Joined {
		if v.DriftingNoteID != 0 {
			var a model.DriftingPicture
			err = mysql.DB.Where("id = ?", v.DriftingNoteID).First(&a).Error
			if err != nil {
				return nil, err
			}
			pictures = append(pictures, a)
		}
	}
	return pictures, nil
}
