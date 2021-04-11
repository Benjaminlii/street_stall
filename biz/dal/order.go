package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
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

// InsertOrder 插入一个order对象
func InsertOrder(insertOrder *model.Order) *model.Order {
	db := GetDB()
	db.Create(insertOrder)
	if err := db.Error; err != nil {
		log.Printf("[service][order][insertOrder] db insert error, err:%s", err)
	}
	return insertOrder
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

// findOrder 根据传入的db查询order
func findOrder(db *gorm.DB) (ans []model.Order) {
	db.Find(&ans)
	return
}
