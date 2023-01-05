package mysql

import (
	"Drifting/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:@tcp(localhost:3306)/drifting?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.AutoMigrate(model.User{}, model.Friend{}, model.AddingFriend{}, model.DriftingNote{}, model.DriftingNovel{}, model.DriftingDrawing{}, model.DriftingPicture{})
	if err != nil {
		panic(err)
		return
	}
}
