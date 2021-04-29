package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
	"time"
)

// GetAllTodayOrderByLocationId 查询某摊位当天的所有order
func GetAllTodayOrderByLocationId(locationId uint) []model.Order {
	db := GetDB()
	db = filterByLocationId(db, locationId)
	db = filterByTodayCreated(db)
	orders := findOrder(db)
	return orders
}

// GetOrderByMerchantIdOrderByCreatedAtDesc 根据merchantId查询其预约单，按照createdAt降序排列
func GetOrderByMerchantIdOrderByCreatedAtDesc(merchantId uint) (orders []model.Order) {
	db := GetDB()
	db = filterByMerchantId(db, merchantId)
	db = orderByCreatedAt(db, true)
	db.Find(&orders)
	return orders
}

// GetOrderByLocationIdNowInUsing 根据摊位id查询当前正在使用的预约单
func GetOrderByLocationIdNowInUsing(locationId uint) *model.Order {
	db := GetDB()
	db = filterByNowUsing(db)
	db = filterByLocationId(db, locationId)
	order := selectOrder(db)
	return order
}

// InsertOrder 插入一个order对象
func InsertOrder(insertOrder *model.Order) *model.Order {
	db := GetDB()
	db.Create(insertOrder)
	if err := db.Error; err != nil {
		log.Printf("[service][order][insertOrder] db insert error, err:%s", err)
		panic(err)
	}
	return insertOrder
}

// GetOrderByMerchantIdOrderByCreatedAtDesc 根据merchantId查询其预约单，按照createdAt降序排列
func GetOrderById(orderId uint) *model.Order {
	db := GetDB()
	db = filterById(db, orderId)
	order := selectOrder(db)
	return order
}

// GetTodayOrderByStatusAndReserveTime 获取当前某个预约时间段内，某个状态的预约单，也可次查询前一天的
func GetTodayOrderByStatusAndReserveTime(status int, reserveTime uint, ifYesterday bool) []model.Order {
	db := GetDB()
	db = filterByStatus(db, status)
	db = filterByReserveTime(db, reserveTime)
	if ifYesterday {
		db = filterByOneDayCreated(db)
	} else {
		db = filterByTodayCreated(db)
	}
	orders := findOrder(db)
	return orders
}

// GetTodayOrderByStatus 获取今天某个状态的所有预约单
func GetTodayOrderByStatus(status int) []model.Order {
	db := GetDB()
	db = filterByStatus(db, status)
	db = filterByTodayCreated(db)
	orders := findOrder(db)
	return orders
}

// SaveOrder 更新并覆盖order
func SaveOrder(order *model.Order) {
	db := GetDB()
	db.Save(order)
}

// filterByLocationId 通过摊位Id过滤
func filterByLocationId(db *gorm.DB, locationId uint) *gorm.DB {
	db = db.Where("location_id = ?", locationId)
	return db
}

// filterByMerchantId 通过商户Id过滤
func filterByMerchantId(db *gorm.DB, merchantId uint) *gorm.DB {
	db = db.Where("merchant_id = ?", merchantId)
	return db
}

// filterByStatus 通过预约单的状态过滤
func filterByStatus(db *gorm.DB, status int) *gorm.DB {
	db = db.Where("status = ?", status)
	return db
}

// filterByReserveTime 通过预约单的预约时间过滤
func filterByReserveTime(db *gorm.DB, reserveTime uint) *gorm.DB {
	db = db.Where("reserve_time = ?", reserveTime)
	return db
}

// filterByNowUsing 过滤当前时间正在被使用的预约单
func filterByNowUsing(db *gorm.DB) *gorm.DB {
	currentTimeHour := time.Now().Hour()
	rightTimeHour := currentTimeHour
	if currentTimeHour%2 != 0 {
		rightTimeHour -= 0
	}
	db = filterByReserveTime(db, uint(rightTimeHour))
	db = filterByStatus(db, constants.ORDER_STATUS_IN_USING)
	return db
}

// findOrder 根据传入的db查询order
func findOrder(db *gorm.DB) (ans []model.Order) {
	db.Find(&ans)
	return
}

// selectOrder 根据db去查询order模型
func selectOrder(db *gorm.DB) *model.Order {
	order := &model.Order{}
	result := db.First(order)
	if err := result.Error; err != nil {
		log.Printf("[service][order][selectOrder] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return order
}
