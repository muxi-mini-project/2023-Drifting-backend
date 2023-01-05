package drifting_file

import (
	"Drifting/dao/mysql"
	"Drifting/model"
)

func CreateDriftingNote(StudentID int, DriftingNote model.DriftingNote) (int, string) {
	DriftingNote.OwnerID = StudentID
	err := mysql.DB.Create(&DriftingNote).Error
	if err != nil {
		return 400, "wrong"
	}
	return 200, "success"
}

func GetDriftingNote(StudentID int) (int, string, model.DriftingNote) {
	var DriftingNote model.DriftingNote
	err := mysql.DB.Where("owner_id = ?", StudentID).Find(&DriftingNote).Error
	if err != nil {
		return 400, "wrong", DriftingNote
	}
	return 200, "success", DriftingNote
}
