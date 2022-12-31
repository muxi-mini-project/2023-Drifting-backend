package controller

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/services/qiniu"
	"fmt"
	"mime/multipart"
)

// GetUserInfo 获取用户信息
func GetUserInfo(StudentId int) (model.User, error) {
	var user model.User
	err := mysql.DB.Where("student_id = ?", int64(StudentId)).First(&user).Error
	return user, err
}

// UpdateUserInfo 更新用户信息(其一,主要对基本信息进行修改)
func UpdateUserInfo(UpdateUser *model.User) (model.User, error) {
	var OldUser model.User
	OldUser.StudentID = UpdateUser.StudentID
	fmt.Println(OldUser)
	err := mysql.DB.Where("student_id = ?", UpdateUser.StudentID).Updates(&UpdateUser).Error
	return *UpdateUser, err
}

// UpdateUserAvatar 更新用户头像
func UpdateUserAvatar(file *multipart.FileHeader, StudentID int) error {
	// 上传到七牛云
	_, url := qiniu.UploadToQiNiu(file, "user_avatar/")
	OldUser, err := GetUserInfo(StudentID)
	if err != nil {
		return err
	}
	NewUser := OldUser
	NewUser.Avatar = url
	err = mysql.DB.Where("student_id = ?", StudentID).Updates(&NewUser).Error
	return err
}
