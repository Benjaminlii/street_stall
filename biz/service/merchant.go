package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dal"
	"street_stall/biz/drivers"
	"street_stall/biz/util"
)

// GetMerchantsInfoByNameAndPlaceId 通过当前时刻商户名称获取商户信息
func GetMerchantsInfoByNameAndPlaceId(c *gin.Context, placeId uint, merchantName string, category uint) map[string]map[string]string {
	// redis中取出当前正在摆摊的商家id，redis中数据来源于打卡
	redisKey := fmt.Sprintf("%s%d", constants.REDIS_CURRENT_ACTIVE_MERCHANT_PRE, placeId)
	merchantIdStrList, err := drivers.RedisClient.SMembers(redisKey).Result()
	if err != nil {
		log.Printf("[service][merchant][GetMerchantsInfoByNameAndPlaceId] get current merchant ids from redis error, err:%s", err)
		panic(errors.SYSTEM_ERROR)
	}
	// 类型转化
	merchantIds := make([]uint, len(merchantIdStrList))
	for _, merchantIdStr := range merchantIdStrList {
		merchantIds = append(merchantIds, util.StringToUInt(merchantIdStr))
	}

	// 查询得到商家
	merchants := dal.FindMerchantByPlaceIdNameAndCategory(placeId, merchantName, merchantIds, category)

	// 组装结果集
	ans := make(map[string]map[string]string, len(merchants))
	for _, merchant := range merchants {
		entity := make(map[string]string, 4)

		entity["name"] = merchant.Name
		entity["category"] = util.UintToCategoryString(merchant.Category)

		merchantStar := float64(merchant.StarSum) / float64(merchant.CommentCount*1.0)
		merchantStarStr := fmt.Sprintf("%.1f", merchantStar)
		entity["star"] = merchantStarStr

		entity["introduction"] = merchant.Introduction

		ans[util.UintToString(merchant.ID)] = entity
	}

	return ans
}
