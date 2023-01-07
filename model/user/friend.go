package user

import (
	"Drifting/dao/mysql"
	"Drifting/model"
)

// SearchAddFriend 查询添加好友记录
func SearchAddFriend(UserId int64, FriendId int64) error {
	var Adding model.AddingFriend
	Adding.AdderID = UserId
	Adding.TargetID = FriendId
	err := mysql.DB.Where(&Adding).Find(&Adding).Error
	return err
}

// SearchFriends 查询是否已添加好友
func SearchFriends(StudentID int64, FriendId int64) error {
	var SearchFriend model.UserAndFriends
	SearchFriend.UserId = StudentID
	SearchFriend.FriendId = FriendId
	return mysql.DB.Where(&SearchFriend).Find(&SearchFriend).Error
}

// AddFriend 添加好友
func AddFriend(UserId int64, FriendId int64) error {
	var Adding model.AddingFriend
	Adding.AdderID = UserId
	Adding.TargetID = FriendId
	return mysql.DB.Create(&Adding).Error
}

// GetFriend 获取好友列表
func GetFriend(StudentID int64) ([]model.UserInfo, error) {
	var FindFriends []model.UserAndFriends
	err := mysql.DB.Where("user_id = ?", StudentID).Find(&FindFriends).Error
	if err != nil {
		return nil, err
	}
	var FriendsInfo []model.UserInfo
	for _, friend := range FindFriends {
		var thefriend model.Friend
		err := mysql.DB.Where("student_id = ?", friend.FriendId).First(&thefriend).Error
		if err != nil {
			return nil, err
		}
		var timelyInfo model.UserInfo
		timelyInfo.Name = thefriend.Name
		timelyInfo.Sex = thefriend.Sex
		timelyInfo.SelfWord = thefriend.SelfWord
		timelyInfo.Avatar = thefriend.Avatar
		FriendsInfo = append(FriendsInfo, timelyInfo)
	}
	return FriendsInfo, nil
}

// GetRequest 获取好友请求
func GetRequest(StudentID int64) ([]model.UserInfo, error) {
	var FindFriends []model.AddingFriend
	err := mysql.DB.Where("target_id = ?", StudentID).Find(&FindFriends).Error
	if err != nil {
		return nil, err
	}
	var FriendsInfo []model.UserInfo
	for _, friend := range FindFriends {
		var thefriend model.Friend
		err := mysql.DB.Where("student_id = ?", friend.AdderID).First(&thefriend).Error
		if err != nil {
			return nil, err
		}
		var timelyInfo model.UserInfo
		timelyInfo.Name = thefriend.Name
		timelyInfo.Sex = thefriend.Sex
		timelyInfo.SelfWord = thefriend.SelfWord
		timelyInfo.Avatar = thefriend.Avatar
		FriendsInfo = append(FriendsInfo, timelyInfo)
	}
	return FriendsInfo, nil
}

// PassRequest 通过好友申请
func PassRequest(StudentID int64, FriendID int64) error {
	var Adding model.AddingFriend
	Adding.AdderID = StudentID
	Adding.TargetID = FriendID
	err := mysql.DB.Find(&Adding).Error
	if err != nil {
		return err
	}
	var NewUserAndFriend model.UserAndFriends
	NewUserAndFriend.UserId = StudentID
	NewUserAndFriend.FriendId = FriendID
	err = mysql.DB.Create(&NewUserAndFriend).Error
	if err != nil {
		return err
	}
	NewUserAndFriend.UserId = FriendID
	NewUserAndFriend.FriendId = StudentID
	err = mysql.DB.Create(&NewUserAndFriend).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Delete(&Adding).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除好友
func Delete(StudentID int64, FriendID int64) (error, error) {
	var UaF model.UserAndFriends
	UaF.UserId = StudentID
	UaF.FriendId = FriendID
	err1 := mysql.DB.Where(&UaF).Delete(&UaF).Error
	UaF.UserId = FriendID
	UaF.FriendId = StudentID
	err2 := mysql.DB.Where(&UaF).Delete(&UaF).Error
	return err1, err2
}
