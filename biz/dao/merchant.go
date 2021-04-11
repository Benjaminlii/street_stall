package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// GetMerchantByUserId 根据userId获取对应的merchant
func GetMerchantByUserId(userId uint) *model.Merchant {
	db := GetDB()
	db = filterByUserId(db, userId)
	merchant := selectMerchant(db)
	return merchant
}

// InsertMerchant 插入一个merchant对象
func InsertMerchant(insertMerchant *model.Merchant) *model.Merchant {
	db := GetDB()
	db.Create(insertMerchant)
	if err := db.Error; err != nil {
		log.Printf("[service][merchant][InsertMerchant] db insert error, err:%s", err)
		panic(err)
	}
	return insertMerchant
}

// SaveMerchant 更新并覆盖merchant
func SaveMerchant(merchant *model.Merchant) {
	db := GetDB()
	db.Save(merchant)
}

// FindMerchantByPlaceIdNameAndCategory 通过placeId，name和类别查询商户
func FindMerchantByPlaceIdNameAndCategory(placeId uint, merchantName string, merchantIds []uint, category uint) []model.Merchant {
	db := GetDB()
	db = filterByPlaceId(db, placeId)
	db = filterByLikeMerchantName(db, merchantName)
	db = filterByInMerchantIDs(db, merchantIds)

	// 如果category为零，那么不进行过滤
	if category != 0 {
		db = filterByCategory(db, category)
	}

	return findMerchant(db)
}

// filterByLikeMerchantName 通过merchantName模糊查询
func filterByLikeMerchantName(db *gorm.DB, merchantName string) *gorm.DB {
	db = db.Where(fmt.Sprintf("merchant_name like '%%%s%%'", merchantName))
	return db
}

// filterByInMerchantIDs 通过merchantId，进行in过滤
func filterByInMerchantIDs(db *gorm.DB, merchantIds []uint) *gorm.DB {
	db = db.Where("name in (?)", merchantIds)
	return db
}

// filterByCategory 通过商家类型进行过滤
func filterByCategory(db *gorm.DB, category uint) *gorm.DB {
	db = db.Where("category = ?", category)
	return db
}

// selectMerchant 查询merchant
func selectMerchant(db *gorm.DB) *model.Merchant {
	merchant := &model.Merchant{}
	db = db.First(merchant)
	if err := db.Error; err != nil {
		log.Printf("[service][merchant][selectMerchant] db select error, err:%s", err)
		panic(err)
	}
	return merchant
}

// findMerchant 根据传入的db查询Merchant
func findMerchant(db *gorm.DB) (merchants []model.Merchant) {
	db.Find(&merchants)
	return
}
