package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
)

// GetPlaceById 根据placeId查询place
func GetPlaceById(placeId uint) *model.Place {
	place := selectPlace(filterById(GetDB(), placeId))
	if place == nil {
		log.Printf("[service][place][GetPlaceById] place is not exist, placeId:%d", placeId)
		panic(errors.NULL_ERROR)
	}
	return place
}

// AllPlace 获取所有区域
func AllPlace() (places []model.Place) {
	db := GetDB()
	db.Find(&places)
	return
}

// selectPlace 根据db去查询place模型
func selectPlace(db *gorm.DB) *model.Place {
	place := &model.Place{}
	result := db.First(place)
	if err := result.Error; err != nil {
		log.Printf("[service][place][selectPlace] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return place
}
