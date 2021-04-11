package service

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/dal"
	"street_stall/biz/domain/dto"
	"street_stall/biz/util"
)

// GetAllTodayReserveByLocation 得到某个摊位下当天的所有预约信息
func GetAllTodayReserveByLocation(c *gin.Context, locationId uint) map[string]int {
	// 查询该摊位当天的所有order
	orders := dal.GetAllTodayOrderByLocationId(locationId)

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
	orders := dal.GetOrderByMerchantIdOrderByCreatedAtDesc(merchant.ID)

	ans := make([]dto.GetOrderDTO, len(orders))
	for _, order := range orders {
		// 根据预约单获取预约单中的摊位信息和区域信息
		place := dal.GetPlaceById(order.PlaceId)
		location := dal.GetLocationById(order.LocationId)
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
