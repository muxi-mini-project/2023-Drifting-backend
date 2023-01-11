package user

import (
	"Drifting/handler"
	"Drifting/model"
	"Drifting/model/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param  Authorization header string true "token"
// @Success 200 {object} model.UserInfo "{"message":"获取成功"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/user/detail [get]
func GetUserDetails(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	//执行获取信息函数
	UserInfo, err := user.GetUserInfo(StudentID)
	if err != nil {
		c.JSON(400, gin.H{"message": "获取用户信息失败"})
		return
	}
	//返回正确信息
	handler.SendGoodResponse(c, "获取成功", UserInfo)
}

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param  User body model.User true "UserInfo"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router api/v1/user/update [put]
func UpdateUserInfo(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	var UpdateUser model.User
	err := c.BindJSON(&UpdateUser)
	UpdateUser.StudentID = StudentID
	fmt.Println(UpdateUser)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	_, err = user.UpdateUserInfo(&UpdateUser)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}

// @Summary 更新用户头像
// @Description 更新用户头像
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token"
// @Param file formData file true "file"
// @Success 200 {object} handler.Response "{"message":"Success"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router api/v1/user/avatar [put]
func UpdateUserAvatar(c *gin.Context) {
	StudentID := c.MustGet("student_id").(int64)
	f, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = user.UpdateUserAvatar(f, StudentID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "成功"})
}
