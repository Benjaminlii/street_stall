package dal

import (
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
