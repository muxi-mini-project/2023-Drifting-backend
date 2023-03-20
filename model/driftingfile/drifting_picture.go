package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/pkg/errno"
	"Drifting/services/qiniu"
	"mime/multipart"
)

// CreateNewDriftingPicture 创建漂流相机
func CreateNewDriftingPicture(NewPicture model.DriftingPicture) (error, uint) {
	//NewPicture.WriterNumber = 1
	err := mysql.DB.Create(&NewPicture).Error
	if err != nil {
		return err, 0
	}
	var FindPicture model.DriftingPicture
	err = mysql.DB.Where(&NewPicture).Find(&FindPicture).Error
	if err != nil {
		return err, 0
	}
	//var NewJoined model.JoinedDrifting
	//NewJoined.StudentID = NewPicture.OwnerID
	//NewJoined.DriftingPictureID = int64(NewPicture.ID)
	//err = JoinDriftingNote(NewJoined)
	//if err != nil {
	//	return err, 0
	//}
	return err, FindPicture.ID
}

// GetDriftingPicture 获取用户创建的漂流相机
func GetDriftingPicture(StudentID int64) ([]model.DriftingPicture, error) {
	var pictures []model.DriftingPicture
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&pictures).Error
	return pictures, err
}

// JoinNewDriftingPicture 参与漂流相机
func JoinNewDriftingPicture(Joining model.JoinedDrifting) error {
	err := mysql.DB.Where(&Joining).First(&Joining).Error
	if err != nil {
		err = mysql.DB.Create(&Joining).Error
		return err
	}
	return errno.ErrDatabase
}

// DrawPicture 创作漂流相机
func DrawPicture(StudentID int64, NewPictureContact model.PictureContact, f *multipart.FileHeader) error {
	NewPictureContact.WriterID = StudentID
	_, url := qiniu.UploadToQiNiu(f, "drifting_picture/")
	NewPictureContact.TheWords = url
	err := mysql.DB.Create(&NewPictureContact).Error
	var NewInfo model.DriftingPicture
	err = mysql.DB.Where("id = ?", NewPictureContact.FileID).Find(&NewInfo).Error
	if err != nil {
		return err
	}
	NewInfo.WriterNumber = NewInfo.WriterNumber + 1
	err = mysql.DB.Where("id = ?", NewPictureContact.FileID).Updates(&NewInfo).Error
	return err
}

// RefusePictureInvite 拒绝漂流相机邀请
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
	Note.SetNumber = Note.SetNumber - 1
	err = mysql.DB.Where("id = ?", Note.ID).Updates(&Note).Error
	return err
}

// DriftingPictureDetail 获取漂流相机内容
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

// GetJoinedDriftingPictures 获取某人加入的漂流相机
func GetJoinedDriftingPictures(StudentID int64) ([]model.DriftingPicture, error) {
	var pictures []model.DriftingPicture
	var Joined []model.JoinedDrifting
	err := mysql.DB.Where("student_id = ?", StudentID).Find(&Joined).Error
	if err != nil {
		return nil, err
	}
	for _, v := range Joined {
		if v.DriftingPictureID != 0 {
			var a model.DriftingPicture
			err = mysql.DB.Where("id = ?", v.DriftingPictureID).Find(&a).Error
			if err != nil {
				return nil, err
			}
			pictures = append(pictures, a)
		}
	}
	return pictures, nil
}

// RandomRecommendPicture 随机推荐漂流相机
func RandomRecommendPicture(StudentID int64) (model.DriftingPicture, error) {
	var pictures []model.DriftingPicture
	err := mysql.DB.Not("kind", 1).Not("number", 1, 0).Not("owner_id", StudentID).Find(&pictures).Error
	if err != nil {
		return model.DriftingPicture{}, err
	}
	m1 := make(map[int]model.DriftingPicture)
	for i := 0; i < len(pictures); i++ {
		m1[i] = pictures[i]
	}
	var ret model.DriftingPicture
	for _, v := range m1 {
		ret = v
		break
	}
	for k, _ := range m1 {
		delete(m1, k)
	}
	return ret, nil
}

// DeletePicture 删除指定漂流相机
func DeletePicture(ThePicture model.DriftingPicture) error {
	err := mysql.DB.Where(&ThePicture).Delete(&ThePicture).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Where("drifting_picture_id = ?", ThePicture.ID).Delete(&model.JoinedDrifting{}).Error
	return err
}
