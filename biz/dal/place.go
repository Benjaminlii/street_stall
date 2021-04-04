package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/domain/model"
)

// AllPlace 获取所有区域
func AllPlace(db *gorm.DB) (places []model.Place) {
	db.Find(&places)
	return
}

// SelectPlace 根据db去查询place模型
func SelectPlace(db *gorm.DB) *model.Place {
	place := &model.Place{}
	db.First(place)
	if err := db.Error; err != nil {
		log.Printf("[service][place][SelectPlace] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(constants.SYSTEM_ERROR)
		}
	}

	return place
}
