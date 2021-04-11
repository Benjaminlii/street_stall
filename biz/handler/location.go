package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// Reserve 预约摊位
func Reserve(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][location][Reserve] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlaceId := param["place_id"]
	locationIdStr, haveLocationId := param["location_id"]
	reserveTimeStr, haveReserveTime := param["reserve_time"]
	comment, haveComment := param["comment"]
	if !(havePlaceId && haveLocationId && haveReserveTime && haveComment) {
		log.Printf("[service][location][Reserve] request type error, err:%s", err)
		panic(errors.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)
	locationId := util.StringToUInt(locationIdStr)
	reserveTime := util.StringToUInt(reserveTimeStr)

	order := service.Reserve(c, placeId, locationId, reserveTime, comment)
	log.Printf("[service][location][Reserve] reserve location success, placeId:%d, locationId:%d, merchantId:%d",
		placeId, locationId, order.MerchantId)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}

// GetMerchantsByPlaceId 通过区域（+分类）获取摊位上的商户信息，当前时刻（游客端）
func GetMerchantsByPlaceId(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][location][GetMerchantsByPlaceId] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlaceId := param["place_id"]
	categoryStr, haveCategory := param["category"]
	if !(havePlaceId && haveCategory) {
		log.Printf("[service][location][GetMerchantsByPlaceId] request type error, err:%s", err)
		panic(errors.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)
	category := util.StringToUInt(categoryStr)

	service.GetMerchantsByPlaceId(c, placeId, category)
	// todo
}
