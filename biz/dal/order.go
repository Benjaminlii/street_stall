package dal

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/domain/model"
)

// FilterByLocationId 通过摊位Id过滤
func FilterByLocationId(db *gorm.DB, locationId uint) *gorm.DB {
	db = db.Where("location_id = ?", locationId)
	return db
}


// FindOrder 根据传入的db查询order
func FindOrder(db *gorm.DB) (ans []model.Order) {
	db.Find(&ans)
	return
}
