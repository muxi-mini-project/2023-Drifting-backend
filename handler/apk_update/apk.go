package apk_update

import (
	"Drifting/handler"
	"Drifting/model/apk"
	"github.com/gin-gonic/gin"
)

// @Summary 上传apk
// @Description 用文件上传apk，不需要写对应功能，建议在apifox上进行更新
// @Tags apk
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param apk formData file true "apk"
// @Param version formData string true "version"
// @Success 200 {object} handler.Response "{"message":"上传成功"}"
// @Failure 400 {object} handler.Response "{"message":"上传失败"}"
// @Router /api/v1/apk/update [post]
func UploadApk(c *gin.Context) {
	f, err := c.FormFile("apk")
	if err != nil {
		handler.SendBadResponse(c, "上传失败", err)
		return
	}
	ver := c.PostForm("version")
	err = apk.Update(f, ver)
	if err != nil {
		handler.SendBadResponse(c, "上传失败", err)
		return
	}
	handler.SendGoodResponse(c, "上传成功", err)
}

// @Summary 获取最新版本号
// @Description 获取最新版本号
// @Tags apk
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.Response "{"message":"获得版本号"}"
// @Failure 400 {object} handler.Response "{"message":"获取版本号失败"}"
// @Router /api/v1/apk/get_version [get]
func GetVersion(c *gin.Context) {
	err, version := apk.GerVersion()
	if err != nil {
		handler.SendBadResponse(c, "获取版本号失败", nil)
		return
	}
	handler.SendGoodResponse(c, "或取版本号成功", version)
}
