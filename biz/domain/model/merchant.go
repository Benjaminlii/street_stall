package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Merchant 商户
type Merchant struct {
	gorm.Model
	base.Row
	UserId       uint   `gorm:"column:user_id"`        // 用户id
	Name         string `gorm:"column:name"`           // 商户名称
	Category     uint   `gorm:"column:category;index"` // 商户分类
	StarSum      uint   `gorm:"column:star_sum"`       // 累计星数
	CommentCount uint   `gorm:"column:comment_count"`  // 累计评价数
	Introduction string `gorm:"column:introduction"`   // 商户简介
}
