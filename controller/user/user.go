package user

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/services/qiniu"
	"fmt"
	"mime/multipart"
)

// GetUserInfo 获取用户信息
func GetUserInfo(StudentId int64) (model.UserInfo, error) {
	var user model.User
	err := mysql.DB.Where("student_id = ?", StudentId).First(&user).Error
	var userInfo model.UserInfo
	userInfo.Name = user.Name
	userInfo.Sex = user.Sex
	userInfo.SelfWord = user.SelfWord
	userInfo.Avatar = user.Avatar
	return userInfo, err
}

// UpdateUserInfo 更新用户信息(主要对基本信息进行修改)
func UpdateUserInfo(UpdateUser *model.User) (model.User, error) {
	var OldUser model.User
	OldUser.StudentID = UpdateUser.StudentID
	fmt.Println(OldUser)
	//更新用户信息
	err := mysql.DB.Where("student_id = ?", UpdateUser.StudentID).Updates(&UpdateUser).Error
	return *UpdateUser, err
}

// UpdateUserAvatar 更新用户头像
func UpdateUserAvatar(file *multipart.FileHeader, StudentID int64) error {
	// 上传到七牛云
	_, url := qiniu.UploadToQiNiu(file, "user_avatar/")
	OldUser, err := GetUserInfo(StudentID)
	if err != nil {
		return err
	}
	NewUser := OldUser
	NewUser.Avatar = url
	err = mysql.DB.Where("student_id = ?", StudentID).Updates(&NewUser).Error
	var FriendsInfo model.Friend
	FriendsInfo.Avatar = url
	mysql.DB.Where("student_id = ?", StudentID).Updates(&FriendsInfo)
	return err
}
