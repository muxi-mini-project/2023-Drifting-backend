package user

import (
	"Drifting/dao/mysql"
	"Drifting/handler"
	"Drifting/model"
	"Drifting/router/middleware"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 登录
// @Tags user
// @Description 一站式登录
// @Accep application/json
// @Produce application/json
// @Param object body model.User true "登录用户信息"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001", "message":"Password or account wrong."} 身份认证失败 重新登录"
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /login [post]
func Login(c *gin.Context) {
	var u model.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if u.StudentID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	pwd := u.PassWord
	//首次登录，验证一站式
	//判断是否首次登录
	result := mysql.DB.Where("student_id = ?", u.StudentID).First(&u)
	if result.Error != nil {
		id := strconv.Itoa(int(u.StudentID))
		_, err := model.GetUserInfoFormOne(id, pwd)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Password or account is wrong.")
			return
		}
		//对用户信息初始化
		u.Name = " "
		//对密码进行base64加密
		u.PassWord = base64.StdEncoding.EncodeToString([]byte(u.PassWord))
		mysql.DB.Create(&u)
	} else {
		//在数据库中解密比较
		password, _ := base64.StdEncoding.DecodeString(u.PassWord)

		if string(password) != pwd {
			c.JSON(http.StatusUnauthorized, "Password or account is wrong.")
			return
		}
	}

	signedToken, err := middleware.GenerateToken(u.StudentID)
	if err != nil {
		log.Println(err)
	}
	handler.SendGoodResponse(c, "将student_id作为token保留", signedToken)
}
