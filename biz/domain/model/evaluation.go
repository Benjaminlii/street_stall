package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Evaluation 游客评价
type Evaluation struct {
	gorm.Model
	base.Row
	Star       uint     `gorm:"column:star"`            // 评价星级
	Content    string   `gorm:"column:content"`         // 评价内容
	MerchantId uint     `gorm:"column:merchant_id"`     // 商户id
	VisitorId  uint     `gorm:"column:visitor_id"`      // 游客id
}
