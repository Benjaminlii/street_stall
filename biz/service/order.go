package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/domain/dto"
	"street_stall/biz/drivers"
	"street_stall/biz/util"
	"time"
)

// GetAllTodayReserveByLocation 得到某个摊位下当天的所有预约信息
func GetAllTodayReserveByLocation(c *gin.Context, locationId uint) map[string]int {
	// 查询该摊位当天的所有order
	orders := dao.GetAllTodayOrderByLocationId(locationId)

	reserveInfoMap := make(map[string]int, len(orders))

	for _, order := range orders {
		reserveInfoMap[util.UintToString(order.ReserveTime)] = order.Status
	}

	return reserveInfoMap
}

// GetOrderByCurrentMerchant 获取当前用户对应的商户其预约单信息，created_at降序排序
func GetOrderByCurrentMerchant(c *gin.Context) []dto.GetOrderDTO {
	// 获取当前登录的商户
	merchant := GetMerchantByCurrentUser(c)
	// 查询其预约单信息
	orders := dao.GetOrderByMerchantIdOrderByCreatedAtDesc(merchant.ID)

	ans := make([]dto.GetOrderDTO, len(orders))
	for _, order := range orders {
		// 根据预约单获取预约单中的摊位信息和区域信息
		place := dao.GetPlaceById(order.PlaceId)
		location := dao.GetLocationById(order.LocationId)
		getOrderDTO := dto.GetOrderDTO{
			OrderId:  order.ID,
			CreateAt: order.CreatedAt.Unix(),
			Status:   order.Status,
			Location: struct {
				Place struct {
					Name string `json:"name"`
				} `json:"place"`
				Number       int    `json:"number"`
				Introduction string `json:"introduction"`
			}{
				Place: struct {
					Name string `json:"name"`
				}{Name: place.Name},
				Number:       location.Number,
				Introduction: location.Introduction,
			},
			ReserveTime: order.ReserveTime,
			Remark:      order.Remark,
		}
		ans = append(ans, getOrderDTO)
	}
	return ans
}

// ClockIn 商户到预约时间打卡使用摊位
func ClockIn(c *gin.Context, orderId uint) {
	// 获取当前登录的商户
	merchant := GetMerchantByCurrentUser(c)
	// 获取预约单
	order := dao.GetOrderById(orderId)

	// 校验商户是否一致
	if merchant.ID != order.ID {
		log.Printf("[service][order][ClockIn] order is not belong current merchant, current merchant name:%s, order id:%d", merchant.Name, order.ID)
		panic(errors.ORDER_MERCHANT_ERROR)
	}

	// 校验时间
	reserveTimeInt := int(order.ReserveTime)
	todayFirstSecond := util.GetTodayFirstSecond()
	currentHour := time.Now().Hour()
	// 校验当前小时大于预约时间，并且小于预约时间+2，并且订单的创建时间是今天
	if !(currentHour > reserveTimeInt &&
		currentHour < reserveTimeInt+2 &&
		order.CreatedAt.After(todayFirstSecond)) {
		log.Printf("[service][order][ClockIn] time is not right")
		panic(errors.ORDER_RESERVE_TIME_ERROR)
	}

	// 进行打卡
	// 更新订单状态
	order.Status = constants.ORDER_STATUS_IN_USING
	dao.SaveOrder(order)
	// 同步redis，将当前商户的id添加到redis中当前地区活跃摆摊的set中
	key := fmt.Sprintf("%s%d", constants.REDIS_CURRENT_ACTIVE_MERCHANT_PRE, order.PlaceId)
	drivers.RedisClient.HSet(key, util.UintToString(merchant.ID), order.LocationId)
}

// QuitOrder 商户退订预约单，即取消预约，需判断时间，预约单状态
func QuitOrder(c *gin.Context, orderId uint) {
	// 获取当前登录的商户
	merchant := GetMerchantByCurrentUser(c)
	// 获取预约单
	order := dao.GetOrderById(orderId)

	// 校验商户是否一致
	if merchant.ID != order.ID {
		log.Printf("[service][order][QuitOrder] order is not belong current merchant, current merchant name:%s, order id:%d", merchant.Name, order.ID)
		panic(errors.ORDER_MERCHANT_ERROR)
	}

	// 校验时间
	reserveTimeInt := int(order.ReserveTime)
	todayFirstSecond := util.GetTodayFirstSecond()
	currentHour := time.Now().Hour()
	// 校验当前小于预约时间，并且订单的创建时间是今天
	if !(currentHour < reserveTimeInt &&
		order.CreatedAt.After(todayFirstSecond)) {
		log.Printf("[service][order][ClockIn] time is not right")
		panic(errors.ORDER_RESERVE_TIME_ERROR)
	}

	// 进行退订
	// 更新订单状态，这里被取消的订单也视为视为过期状态
	order.Status = constants.ORDER_STATUS_EXPIRED
	dao.SaveOrder(order)
}
