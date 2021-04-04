package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Visitor 游客
type Visitor struct {
	gorm.Model
	base.Row
	UserId       uint   `gorm:"column:user_id"`      // 用户id
	Name         string `gorm:"column:name"`         // 游客昵称
	Introduction string `gorm:"column:introduction"` // 个人简介
}
