package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// GetOrders 根据当前用户对应的商户信息查询其预约单
func GetOrders(c *gin.Context) {
	defer util.SetResponse(c)

	getOrderDTOs := service.GetOrderByCurrentMerchant(c)

	c.Set(constants.DATA, getOrderDTOs)
}

// ClockIn 商户到预约时间打卡使用摊位
func ClockIn(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][order][ClockIn] request type error, err:%s", err)
		panic(err)
	}
	OrderIdStr, haveOrderId := param["order_id"]
	if !haveOrderId {
		log.Printf("[service][order][ClockIn] request type error, err:%s", err)
		panic(errors.REQUEST_TYPE_ERROR)
	}
	orderId := util.StringToUInt(OrderIdStr)

	service.ClockIn(c, orderId)

	log.Printf("[service][order][ClockIn] clock in success, merchant order id:%d", orderId)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}

// QuitOrder 商户退订预约单
func QuitOrder(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][order][QuitOrder] request type error, err:%s", err)
		panic(err)
	}
	OrderIdStr, haveOrderId := param["order_id"]
	if !haveOrderId {
		log.Printf("[service][order][QuitOrder] orderId is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	orderId := util.StringToUInt(OrderIdStr)

	service.QuitOrder(c, orderId)

	log.Printf("[service][order][QuitOrder] quit order success, merchant order id:%d", orderId)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}
