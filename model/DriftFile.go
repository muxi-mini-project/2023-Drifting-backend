package model

import "gorm.io/gorm"

type DriftingNote struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Number  int              `json:"number"`
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-notes"`
}

type DriftingNovel struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-novel"`
}

type DriftingDrawing struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-drawing"`
}

type DriftingPicture struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Contact string `json:"contact" gorm:"size:255"`
	Cover   string `json:"cover" gorm:"size:255"`
	OwnerID int64
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-picture"`
}

type NoteContact struct {
	FileID   int64  `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	Text     string `json:"text" gorm:"type:text"`
}

type NoteInfo struct {
	Name     string
	OwnerID  int64
	Contacts []NoteContact `json:"contacts"`
}
