package handler

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/constants"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// GetOrders 根据当前用户对应的商户信息查询其预约单
func GetOrders(c *gin.Context) {
	defer util.SetResponse(c)

	getOrderDTOs := service.GetOrderByCurrentMerchant(c)

	c.Set(constants.DATA, getOrderDTOs)
}
