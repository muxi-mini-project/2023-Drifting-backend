package model

import "gorm.io/gorm"

type Apk struct {
	gorm.Model
	Addr    string
	Version string
}
