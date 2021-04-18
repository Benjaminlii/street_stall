package dao

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/constants/errors"
	"street_stall/biz/drivers"
	"street_stall/biz/util"
)

func GetDB() *gorm.DB {
	return drivers.DB
}

// filterByTodayCreated 按照时间过滤今天创建的
func filterByTodayCreated(db *gorm.DB) *gorm.DB {
	todayFirstSecond := util.GetTodayFirstSecond()
	todayLastSecond := util.GetTodayLastSecond()

	db = db.Where("created_at between ? and ?", todayFirstSecond, todayLastSecond)
	return db
}

// filterByOneDayCreated 按照时间过滤某一天创建的
func filterByOneDayCreated(db *gorm.DB) *gorm.DB {
	yesterday := util.GetYesterdayTime()
	oneDayFirstSecond := util.GetOneDayFirstSecond(yesterday)
	oneDayLastSecond := util.GetOneDayLastSecond(yesterday)

	db = db.Where("created_at between ? and ?", oneDayFirstSecond, oneDayLastSecond)
	return db
}

// filterById 按照id过滤
func filterById(db *gorm.DB, id uint) *gorm.DB {
	db = db.Where("id = ?", id)
	return db
}

// orderByCreatedAt 按照创建时间排序
// isDesc:是否降序
func orderByCreatedAt(db *gorm.DB, isDesc bool) *gorm.DB {
	if isDesc {
		db = db.Order("created_at DESC")
	} else {
		db = db.Order("created_at ASC")
	}
	return db
}

// limit 分页查询
func limit(db *gorm.DB, offset uint, count uint) *gorm.DB {
	if offset > 0 && count > 0 {
		db = db.Limit(count).Offset((offset - 1) * count)
	} else {
		panic(errors.DB_LIMIT_ERROR)
	}
	return db
}
