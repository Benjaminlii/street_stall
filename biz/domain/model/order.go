package model

import (
	"gorm.io/gorm"
	"street_stall/biz/domain/model/base"
)

// Order 预约单
type Order struct {
	gorm.Model
	base.Row
	Status      int    `gorm:"column:status"`       // 当前状态
	PlaceId     uint   `gorm:"column:place_id"`     // 区域id
	LocationId  uint   `gorm:"column:location_id"`  // 摊位id
	MerchantId  uint   `gorm:"column:merchant_id"`  // 商户id
	ReserveTime uint   `gorm:"column:reserve_time"` // 预约时间，8,10,12,14,16,18,20,22
	Remark      string `gorm:"column:remark"`       // 备注
}
