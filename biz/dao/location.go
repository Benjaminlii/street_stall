package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// GetLocationsByPlaceId 获取某个place（区域）下的所有Location（摊位）
func GetLocationsByPlaceId(placeId uint) []model.Location {
	db := GetDB()
	db = filterByPlaceId(db, placeId)
	locations := findLocation(db)
	return locations
}

// GetLocationById 根据id获取location
func GetLocationById(locationId uint) *model.Location {
	db := GetDB()
	db = filterById(db, locationId)
	location := selectLocation(db)
	return location
}

// findLocation 根据传入的db查询Location
func findLocation(db *gorm.DB) (locations []model.Location) {
	db.Find(&locations)
	return
}

// selectLocation 查询location
func selectLocation(db *gorm.DB) *model.Location {
	location := &model.Location{}
	result := db.First(location)
	if err := result.Error; err != nil {
		log.Printf("[service][location][selectLocation] db select error, err:%s", err)
		panic(err)
	}
	return location
}

// filterByPlaceId 通过placeId过滤
func filterByPlaceId(db *gorm.DB, placeId uint) *gorm.DB {
	db = db.Where("place_id = ?", placeId)
	return db
}
