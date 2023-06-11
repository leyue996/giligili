package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title    string //标签
	Info     string
	BossId   uint
	BossName string
}
