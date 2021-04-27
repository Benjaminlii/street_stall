package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Location 摊位
type Location struct {
	gorm.Model
	base.Row
	PlaceId      uint   `gorm:"column:place_id"`     // 区域id
	Number       int    `gorm:"column:number"`       // 摊位编号(在某区域内)
	Status       int    `gorm:"column:status;index"` // 摊位状态
	Area         int    `gorm:"column:area"`         // 摊位面积
	Introduction string `gorm:"column:introduction"` // 摊位简介
}
