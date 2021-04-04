package dal

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/domain/model"
)

// FindLocation 根据传入的db查询Location
func FindLocation(db *gorm.DB) (locations []model.Location) {
	db.Find(&locations)
	return
}

// FilterByPlaceId 通过placeId过滤
func FilterByPlaceId(db *gorm.DB, placeId uint) *gorm.DB {
	db = db.Where("place_id = ?", placeId)
	return db
}
