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
		log.Printf("[service][evaluation][DoEvaluation] has nil in start, merchantId and content")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	start := util.StringToUInt(startStr)
	merchantId := util.StringToUInt(merchantIdStr)

	service.DoEvaluation(c, merchantId, start, content)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}

// GetEvaluationsByMerchantId 根据商户id获取的评价信息
// 暂时废弃
func GetEvaluationsByMerchantId(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][evaluation][GetEvaluationsByMerchantId] request type error, err:%s", err)
		panic(err)
	}
	merchantIdStr, haveMerchantId := param["merchant_id"]
	if !haveMerchantId {
		log.Printf("[service][evaluation][GetEvaluationsByMerchantId] merchantId is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	merchantId := util.StringToUInt(merchantIdStr)

	ans := service.GetEvaluationsByMerchantId(c, merchantId)

	// 设置请求响应
	c.Set(constants.DATA, ans)
}
