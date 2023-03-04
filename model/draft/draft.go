package draft

import (
	"Drifting/dao/mysql"
	"Drifting/model"
)

// CreateDraft 创建草稿箱
func CreateDraft(StudentID int64, NewDraft model.Draft) error {
	NewDraft.OwnerID = StudentID
	err := mysql.DB.Create(&NewDraft).Error
	return err
}

// WriteDraft 写草稿
func WriteDraft(StudentID int64, TheContact model.DraftContact) error {
	TheContact.WriterID = StudentID
	err := mysql.DB.Create(&TheContact).Error
	return err
}

// GetDrafts 获取某人的草稿箱
func GetDrafts(StudentID int64) ([]model.Draft, error) {
	var drafts []model.Draft
	err := mysql.DB.Where("owner_id=?", StudentID).Find(&drafts).Error
	return drafts, err
}
