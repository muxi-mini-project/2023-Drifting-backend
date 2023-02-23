package model

type UserAndFriends struct {
	UserId   int64 `json:"userId"`
	FriendId int64 `json:"friendId"`
}

func (UserAndFriends) TableName() string {
	return "user_friends"
}
