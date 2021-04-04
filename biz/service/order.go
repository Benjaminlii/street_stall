package service

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/dal"
	"street_stall/biz/util"
)

// GetAllTodayReserveByLocation 得到某个摊位下当天的所有预约信息
func GetAllTodayReserveByLocation(c *gin.Context, locationId uint) map[string]int {
	// 查询该摊位当天的所有order
	db := dal.GetDB()
	db = dal.FilterByLocationId(db, locationId)
	db = dal.FilterByTodayCreated(db)
	orders := dal.FindOrder(db)

	reserveInfoMap := make(map[string]int, len(orders))

	for _, order := range orders {
		reserveInfoMap[util.UintToString(order.ReserveTime)] = order.Status
	}

	return reserveInfoMap
}
