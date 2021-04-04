package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/domain/model"
)

// GetPlaceById 根据placeId查询place
func GetPlaceById(placeId uint) *model.Place {
	place := selectPlace(filterById(GetDB(), placeId))
	if place == nil{
		log.Printf("[service][place][GetPlaceById] place is not exist, placeId:%d", placeId)
		panic(constants.NULL_ERROR)
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
	db.First(place)
	if err := db.Error; err != nil {
		log.Printf("[service][place][selectPlace] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(constants.SYSTEM_ERROR)
		}
	}

	return place
}
