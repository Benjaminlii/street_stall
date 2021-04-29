package model

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/domain/model/base"
)

type Administrator struct {
	gorm.Model
	base.Row
	Username string `gorm:"column:username;unique_index"` // 管理员用户名
	Password string `gorm:"column:password"`              // 密码
}
