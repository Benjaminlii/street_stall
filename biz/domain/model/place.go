package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Place 区域
type Place struct {
	gorm.Model
	base.Row
	Name string `gorm:"column:name"` // 区域名称
}
