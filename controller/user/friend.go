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
	err := mysql.DB.Where("user_id").Find(&FindFriends).Error
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
