package dal

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/drivers"
	"time"
)

func GetDB() *gorm.DB {
	return drivers.DB
}


// FilterByTodayCreated 按照时间过滤今天创建的
func FilterByTodayCreated(db *gorm.DB) *gorm.DB {
	currentTime := time.Now()
	todayFirstSecond := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	todayLastSecond := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()

	db = db.Where("created_at between ? and ?", todayFirstSecond, todayLastSecond)
	return db
}
// FilterById 按照id过滤
func FilterById(db *gorm.DB, id uint) *gorm.DB {
	db = db.Where("id = ?", id)
	return db
}