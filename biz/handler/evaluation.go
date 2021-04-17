package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// DoEvaluation 用户提交评价
func DoEvaluation(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][evaluation][DoEvaluation] request type error, err:%s", err)
		panic(err)
	}
	startStr, haveStart := param["start"]
	merchantIdStr, haveMerchantId := param["merchant_id"]
	content, haveContent := param["content"]
	if !(haveStart && haveMerchantId && haveContent) {
		log.Printf("[service][evaluation][DoEvaluation] request type error, err:%s", err)
		panic(errors.REQUEST_TYPE_ERROR)
	}
	start := util.StringToUInt(startStr)
	merchantId := util.StringToUInt(merchantIdStr)

	service.DoEvaluation(c, merchantId, start, content)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}
