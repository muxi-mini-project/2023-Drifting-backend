package driftingfile

import (
	"Drifting/dao/mysql"
	"Drifting/model"
)

// CreateDriftingNote 创建漂流本
func CreateDriftingNote(StudentID int64, DriftingNote model.DriftingNote) error {
	DriftingNote.OwnerID = StudentID
	err := mysql.DB.Create(&DriftingNote).Error
	return err
}

// WriteDriftingNote 参与创作
func WriteDriftingNote(StudentID int64, TheContact model.NoteContact) error {
	TheContact.WriterID = StudentID
	err := mysql.DB.Create(&TheContact).Error
	return err
}
