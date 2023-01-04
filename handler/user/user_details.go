package user

import (
	"Drifting/controller"
	"Drifting/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param  student_id body int true "student_id"
// @Success 200 {object} model.User
// @Router /api/v1/user/detail [get]
func GetUserDetails(c *gin.Context) {
	//从表单中获取学号
	StudentID := c.PostForm("student_id")
	//格式转化进行搜索
	id, err := strconv.Atoi(StudentID)
	if err != nil {
		fmt.Println(err.Error())
	}
	//执行获取信息函数
	UserInfo, err := controller.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "获取用户信息失败"})
		return
	}
	//返回正确信息
	c.JSON(http.StatusOK, UserInfo)
}

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param  User body model.User true "UserInfo"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failure"
// @Router api/v1/user/update [put]
func UpdateUserInfo(c *gin.Context) {
	//StudentID := c.PostForm("student_id")
	//id, err1 := strconv.Atoi(StudentID)
	var user model.User
	err := c.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	_, err = controller.UpdateUserInfo(&user)
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
// @Param file formData file true "file"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failure"
// @Router api/v1/user/avatar [put]
func UpdateUserAvatar(c *gin.Context) {
	f, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	StudentID := c.PostForm("student_id")
	id, err := strconv.Atoi(StudentID)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = controller.UpdateUserAvatar(f, id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "成功"})
}
