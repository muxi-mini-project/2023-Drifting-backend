package model

import "gorm.io/gorm"

type DriftingNote struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Number  int    `json:"number"`
	Kind    string `json:"kind" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-notes"`
}

type DriftingNovel struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Kind    string `json:"kind" gorm:"size:255" binding:"required"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-novel"`
}

type DriftingDrawing struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Number  int    `json:"number" gorm:"size:255"`
	Kind    string `json:"kind" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-drawing"`
}

type DriftingPicture struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Kind    string `json:"kind" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-picture"`
}

type NoteContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	Text     string `json:"text" gorm:"type:text"`
}

type Invite struct {
	HostID   int64  `json:"host_id"`
	FriendID int64  `json:"friend_id"`
	FileID   int64  `json:"file_id" binding:"required"`
	FileKind string `json:"file_kind" gorm:"size:255" binding:"required"`
}

type NoteInfo struct {
	Name     string
	OwnerID  int64
	Contacts []NoteContact `json:"contacts"`
}

type DrawingContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	Picture  string `json:"picture" grom:"size:255"`
}

type DrawingInfo struct {
	Name     string
	OwnerID  int64
	Contacts []DrawingContact
}
