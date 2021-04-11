package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/domain/model"
)

// GetLocationsByPlaceId 获取某个place（区域）下的所有Location（摊位）
func GetLocationsByPlaceId(c *gin.Context, placeId uint) []model.Location {
	place := dao.GetPlaceById(placeId)
	if place == nil {
		log.Printf("[service][place][GetLocationsByPlaceId] place is not exist")
		panic(errors.NULL_ERROR)
	}
	locations := dao.GetLocationsByPlaceId(placeId)
	return locations
}

// Reserve 预约摊位
func Reserve(c *gin.Context, placeId uint, locationId uint, reserveTime uint, comment string) *model.Order {
	// 获取商户信息
	merchant := GetMerchantByCurrentUser(c)
	place := dao.GetPlaceById(placeId)
	location := dao.GetLocationById(locationId)

	order := &model.Order{
		Status:      constants.ORDER_STATUS_TO_BE_USED,
		PlaceId:     place.ID,
		LocationId:  location.ID,
		MerchantId:  merchant.ID,
		ReserveTime: reserveTime,
		Remark:      comment,
	}

	insertOrder := dao.InsertOrder(order)

	return insertOrder
}
