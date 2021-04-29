package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// User 用户
type User struct {
	gorm.Model
	base.Row
	Username     string `gorm:"column:username;unique_index"` // 用户名
	Password     string `gorm:"column:password"`              // 密码
	UserIdentity uint   `gorm:"column:user_identity"`         // 用户身份（1：商户，2：游客，3：政府）
}
