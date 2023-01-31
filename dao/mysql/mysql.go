package mysql

import (
	"Drifting/model"
	"Drifting/services/parseyaml"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	v := parseyaml.GetYaml()
	username := v.GetString("db.username")
	password := v.GetString("db.password")
	addr := v.GetString("db.addr")
	port := v.GetInt("db.port")
	dbname := v.GetString("db.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.AutoMigrate(model.User{}, model.Friend{}, model.Invite{}, model.UserAndFriends{}, model.AddingFriend{}, model.JoinedDrifting{}, model.DriftingNote{}, model.NoteContact{}, model.DriftingNovel{}, model.DriftingDrawing{}, model.DrawingContact{}, model.DriftingPicture{})
	if err != nil {
		panic(err)
		return
	}
}
