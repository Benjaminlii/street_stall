package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sort"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/domain/dto"
	"street_stall/biz/domain/model"
	"street_stall/biz/drivers"
	"street_stall/biz/util"
)

// GetMerchantsInfoByNameAndPlaceId 通过当前时刻商户名称获取商户信息
func GetMerchantsInfoByNameAndPlaceId(c *gin.Context, placeId uint, merchantName string, category uint) map[string]map[string]string {
	// redis中取出当前正在摆摊的商家id，redis中数据来源于打卡
	redisKey := fmt.Sprintf("%s%d", constants.REDIS_CURRENT_ACTIVE_MERCHANT_PRE, placeId)
	merchantIdStrList, err := drivers.RedisClient.HKeys(redisKey).Result()
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
	merchants := dao.FindMerchantByPlaceIdNameAndCategory(placeId, merchantName, merchantIds, category)

	// 组装结果集
	ans := make(map[string]map[string]string, len(merchants))
	for _, merchant := range merchants {
		entity := make(map[string]string, 4)

		entity["name"] = merchant.Name
		entity["category"] = util.UintToCategoryString(merchant.Category)

		merchantStar := merchant.GetStar()
		merchantStarStr := fmt.Sprintf("%.1f", merchantStar)
		entity["star"] = merchantStarStr

		entity["introduction"] = merchant.Introduction

		ans[util.UintToString(merchant.ID)] = entity
	}

	return ans
}

// GetMerchantByMerchantId 根据id获取商户的基础信息，包括商户名称，商户分类，星级评价，商户简介
func GetMerchantByMerchantId(c *gin.Context, merchantId uint) map[string]string {
	ans := make(map[string]string, 4)

	merchant := dao.GetMerchantById(merchantId)

	ans["name"] = merchant.Name
	ans["category"] = util.UintToCategoryString(merchant.Category)
	ans["star"] = fmt.Sprintf("%f", merchant.GetStar())
	ans["introduction"] = merchant.Introduction

	return ans
}

// GetMerchantByPlaceIdAndNumber 根据区域和偏移量获取当前位置上商户的基础信息，包括商户名称，商户分类，星级评价，商户简介
func GetMerchantByPlaceIdAndNumber(c *gin.Context, placeId uint, numberOfPlace uint) map[string]string {
	// 根据区域和偏移量获取摊位
	location := dao.GetLocationByPlaceIdAndNumber(placeId, numberOfPlace)
	// 根据摊位和当前时刻使用确定预约单
	nowUsingOrder := dao.GetOrderByLocationIdNowInUsing(location.ID)
	// 根据预约单得到对应的商户信息
	return GetMerchantByMerchantId(c, nowUsingOrder.MerchantId)
}

// GetMerchantByCurrentUser 获取当前用户对应的商户
func GetMerchantByCurrentUser(c *gin.Context) *model.Merchant {
	// 获取user
	currentUser := util.GetCurrentUser(c)

	if currentUser.UserIdentity != constants.USERIDENTITY_MERCHANT {
		log.Printf("[service][merchant][GetVisitorByCurrentUser] current user is not a merchant")
		panic(errors.AUTHORITY_ERROR)
	}

	merchant := dao.GetMerchantByUserId(currentUser.ID)

	return merchant
}

// GetMerchantsByPlaceId 根据区域id获取商家信息列表（按照星级降序排列，固定数量  15）
func GetMerchantsByPlaceId(c *gin.Context, placeId uint, category uint, isOrderByStart uint, offset uint, count uint) []dto.GetMerchantsDTO {
	// redis中取出当前正在摆摊的商家id，redis中数据来源于打卡
	redisKey := fmt.Sprintf("%s%d", constants.REDIS_CURRENT_ACTIVE_MERCHANT_PRE, placeId)
	merchantIdStrList, err := drivers.RedisClient.HKeys(redisKey).Result()
	if err != nil {
		log.Printf("[service][merchant][GetMerchantsByPlaceId] get current merchant ids from redis error, err:%s", err)
		panic(errors.SYSTEM_ERROR)
	}
	// 类型转化
	merchantIds := make([]uint, len(merchantIdStrList))
	for _, merchantIdStr := range merchantIdStrList {
		merchantIds = append(merchantIds, util.StringToUInt(merchantIdStr))
	}
	ans := make([]dto.GetMerchantsDTO, len(merchantIds))

	merchants := dao.FindMerchantByIdsCategoryLimit(merchantIds, category, offset, count)
	// 这里只能手动排序了

	for _, merchant := range merchants {
		locationId, err := drivers.RedisClient.HGet(redisKey, util.UintToString(merchant.ID)).Result()
		if err != nil {
			log.Printf("[service][merchant][GetMerchantsByPlaceId] get location id from redis error, err:%s", err)
			panic(errors.SYSTEM_ERROR)
		}
		location := dao.GetLocationById(util.StringToUInt(locationId))
		getMerchantsDTO := dto.GetMerchantsDTO{
			MerchantId:   merchant.ID,
			Stars:        merchant.GetStar(),
			Introduction: merchant.Introduction,
			Location: struct {
				LocationId    uint `json:"location_id"`
				NumberOfPlace int  `json:"number_of_place"`
			}{LocationId: location.ID, NumberOfPlace: location.Number},
		}
		ans = append(ans, getMerchantsDTO)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Stars > ans[j].Stars
	})
	return ans
}
