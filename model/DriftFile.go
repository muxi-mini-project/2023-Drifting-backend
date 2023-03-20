package model

import (
	"gorm.io/gorm"
	"time"
)

type DriftingNote struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	Cover        string `json:"cover" gorm:"size:255"`
	OwnerID      int64
	SetNumber    int    `json:"number"`
	WriterNumber int    `json:"writer_number"`
	Kind         int    `json:"kind"`
	Theme        string `json:"theme" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-notes"`
}

type DriftingNovel struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	Cover        string `json:"cover" gorm:"size:255"`
	OwnerID      int64
	SetNumber    int    `json:"number"`
	WriterNumber int    `json:"writer_number"`
	Kind         int    `json:"kind"`
	Theme        string `json:"theme" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-novel"`
}

type DriftingDrawing struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	Cover        string `json:"cover" gorm:"size:255"`
	OwnerID      int64
	SetNumber    int    `json:"number"`
	WriterNumber int    `json:"writer_number"`
	Kind         int    `json:"kind"`
	Theme        string `json:"theme" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-drawing"`
}

type DriftingPicture struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	Cover        string `json:"cover" gorm:"size:255"`
	OwnerID      int64
	SetNumber    int    `json:"number"`
	WriterNumber int    `json:"writer_number"`
	Kind         int    `json:"kind"`
	Theme        string `json:"theme" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-picture"`
}

type Draft struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	Contact  string `json:"contact" gorm:"size:255"`
	Cover    string `json:"cover" gorm:"size:255"`
	OwnerID  int64
	FileKind string `json:"file_kind"`
	Kind     int    `json:"kind"`
	Theme    string `json:"theme" gorm:"size:255"`
	//Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-novel"`
}

type InviteInfo struct {
	Name     string    `json:"name"`
	FileID   uint      `json:"file_id"`
	CreateAt time.Time `json:"createdAt"`
	FileKind string    `json:"fileKind"`
	OwnerID  int64     `json:"owner_id"`
	Cover    string    `json:"cover"`
	Kind     int       `json:"kind"`
	Theme    string    `json:"theme"`
	Number   int       `json:"number"`
}

type NoteContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	TheWords string `json:"the_words" gorm:"size:255"`
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
	Contacts []NoteContact
}

type DrawingContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	TheWords string `json:"the_words" gorm:"size:255"`
}

type DrawingInfo struct {
	Name     string
	OwnerID  int64
	Contacts []DrawingContact
}

type GetFileId struct {
	ID uint `json:"id"`
}

type CreateFile struct {
	Name      string `json:"name" gorm:"size:255"`
	Cover     string `json:"cover" gorm:"size:255"`
	SetNumber int    `json:"number"`
	Kind      string `json:"kind" gorm:"size:255"`
}

type NovelContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	TheWords string `json:"the_words" gorm:"size:255"`
}

type NovelInfo struct {
	Name     string
	OwnerID  int64
	Contacts []NovelContact
}

type PictureContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	TheWords string `json:"the_words" gorm:"size:255"`
}

type PictureInfo struct {
	Name     string
	OwnerID  int64
	Contacts []PictureContact
}

type DraftContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	TheWords string `json:"the_words" gorm:"size:255"`
}

type DraftInfo struct {
	Name     string
	OwnerID  int64
	Contacts []DraftContact `json:"contacts"`
}
