package file_state

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/model/user"
	"Drifting/pkg/errno"
)

// LockOn 上锁
func LockOn(LockRequest model.State) (error, model.State) {
	var TestModel model.State
	err := mysql.DB.Where("file_id = ? AND file_kind = ?", LockRequest.FileID, LockRequest.FileKind).Find(&TestModel).Error
	if err == nil {
		return errno.ErrState, TestModel
	} else {
		err = mysql.DB.Create(&LockRequest).Error
		return err, model.State{}
	}
}

// UnLock 解锁
func UnLock(UnlockRequest model.State) error {
	//var TestModel model.State
	err := mysql.DB.Where(&UnlockRequest).Error
	if err != nil {
		return errno.ErrUnLock
	}
	err = mysql.DB.Delete(&UnlockRequest).Error
	return err
}

// GetLock 获取当前上锁人
func GetLock(GetLockRequest model.State) (error, model.UserInfo) {
	var TestModel model.State
	err := mysql.DB.Where(&GetLockRequest).Find(&TestModel).Error
	if err != nil {
		return err, model.UserInfo{}
	}
	ThisUserInfo, err := user.GetUserInfo(TestModel.WriterID)
	if err != nil {
		return err, model.UserInfo{}
	}
	return err, ThisUserInfo
}
