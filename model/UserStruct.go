package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	StudentID int64  `json:"studentID"`
	PassWord  string `json:"passWord" gorm:"size:255"`
	Name      string `json:"name" gorm:"size:255"`
	Sex       string `json:"sex" gorm:"size:255"`
	Avatar    string `json:"avatar" gorm:"size:255"`
	SelfWord  string `json:"selfWord" gorm:"size:255"`
	//Friends   []Friend `gorm:"many2many:user_friends"`
}

func (u User) Error() string {
	//TODO implement me
	panic("implement me")
}

type AddingFriend struct {
	AdderID  int64 `json:"adderID" `
	TargetID int64 `json:"targetID"`
}

type Friend struct {
	gorm.Model
	StudentID int64  `json:"studentID"`
	Name      string `json:"name" gorm:"size:255"`
	Sex       string `json:"sex" gorm:"size:255"`
	Avatar    string `json:"avatar" gorm:"size:255"`
	SelfWord  string `json:"selfWord" gorm:"size:255"`
	//Friends   []User `gorm:"many2many:user_friends"`
}

type OwnDrifting struct {
	gorm.Model
	StudentID int64 `json:"studentID"`
}

type JoinedDrifting struct {
	gorm.Model
	StudentID        int64             `json:"studentID"`
	DriftingNotes    []DriftingNote    `gorm:"many2many:joined-drifting_drifting-note"`
	DriftingDrawings []DriftingDrawing `gorm:"many2many:joined-drifting_drifting-drawing"`
	DriftingNovels   []DriftingNovel   `gorm:"many2many:joined-drifting_drifting-novel"`
	DriftingPictures []DriftingPicture `gorm:"many2many:joined-drifting_drifting-picture"`
}

type UserInfo struct {
	Name     string
	Sex      string
	SelfWord string
	Avatar   string
}
