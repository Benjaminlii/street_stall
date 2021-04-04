package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// InsertMerchant 插入一个merchant对象
func InsertMerchant(db *gorm.DB, insertMerchant *model.Merchant) *model.Merchant {
	db.Create(insertMerchant)
	if err := db.Error; err != nil {
		log.Printf("[service][merchant][InsertMerchant] db insert error, err:%s", err)
		panic(err)
	}
	return insertMerchant
}

// SelectMerchant 查询merchant
func SelectMerchant(db *gorm.DB) *model.Merchant {
	merchant := &model.Merchant{}
	db = db.First(merchant)
	if err := db.Error; err != nil {
		log.Printf("[service][merchant][SelectMerchant] db select error, err:%s", err)
		panic(err)
	}
	return merchant
}

// SaveMerchant 更新并覆盖merchant
func SaveMerchant(db *gorm.DB, merchant *model.Merchant) {
	db.Save(merchant)
}
