package main

import (
	"Drifting/dao/mysql"
	"Drifting/router"
)

// @title Drifting API
// @version 1.0
// @description 漂流API
// @termsOfService http://swagger.io/terms/
// @contact.name KitZhangYs
// @contact.email SJMbaiyang@163.com
// @host 116.204.121.9:8088
// @BasePath /api/v1
func main() {
	mysql.InitMysql()
	e := router.RouterInit()
	err := e.Run(":8088")
	if err != nil {
		panic(err)
		return
	}
}
