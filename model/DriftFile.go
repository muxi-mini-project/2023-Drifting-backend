package model

import "gorm.io/gorm"

type DriftingNote struct {
	gorm.Model
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Cover   string `json:"cover"`
	OwnerID int
	Owner   OwnDrifting      `gorm:"foreignkey:OwnerID"`
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-notes"`
}

type DriftingNovel struct {
	gorm.Model
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Cover   string `json:"cover"`
	OwnerID int
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-novel"`
}

type DriftingDrawing struct {
	gorm.Model
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Cover   string `json:"cover"`
	OwnerID int
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-drawing"`
}

type DriftingPicture struct {
	gorm.Model
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Cover   string `json:"cover"`
	OwnerID int
	Writers []JoinedDrifting `gorm:"many2many:joined-drifting_drifting-picture"`
}
