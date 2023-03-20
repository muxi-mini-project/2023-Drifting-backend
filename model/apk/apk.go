package apk

import (
	"Drifting/dao/mysql"
	"Drifting/model"
	"Drifting/services/qiniu"
	"fmt"
	"mime/multipart"
)

// Update 更新apk
func Update(apk *multipart.FileHeader, ver string) error {
	apk.Filename = fmt.Sprintf("Drifting-%s.apk", ver)
	_, url := qiniu.UploadToQiNiu(apk, "/apks")
	var ThisApk model.Apk
	ThisApk = model.Apk{
		Addr:    url,
		Version: ver,
	}
	var OldApk model.Apk
	err := mysql.DB.Model(&OldApk).First(&OldApk).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Where(&OldApk).Updates(&ThisApk).Error
	return err
}

func GerVersion() (error, model.Apk) {
	var GetApk model.Apk
	err := mysql.DB.Model(&GetApk).First(&GetApk).Error
	return err, GetApk
}
