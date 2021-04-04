package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// GetPlaceNameToIdMap 获得key-value为place.name-place.id的map，用作当前区域的选择
func GetPlaceNameToIdMap(c *gin.Context) {
	defer util.SetResponse(c)

	// 更新商户信息
	placeNameToIdMap := service.GetPlaceNameToIdMap(c)

	// 设置请求响应
	respMap := map[string]interface{}{
		"place_name_id_map": placeNameToIdMap,
	}

	c.Set(constants.DATA, respMap)
}

// GetLocationMap 获取某地区的摊位map
func GetLocationMap(c *gin.Context){
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][place][GetLocationMap] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlaceId := param["placeId"]
	if !havePlaceId {
		log.Print("[service][question][SubmitQuestion] there is no place id")
		panic(constants.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)

	// 进行业务逻辑计算
	data := service.GetLocationMapAndPlaceInfo(c, placeId)

	// 设置结果集
	c.Set(constants.DATA, data)
}
