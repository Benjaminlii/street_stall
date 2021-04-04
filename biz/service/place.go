package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"street_stall/biz/constants"
	"street_stall/biz/dal"
	"street_stall/biz/util"
	"time"
)

// GetPlaceNameToIdMap 获得key-value为place.name-place.id的map
func GetPlaceNameToIdMap(c *gin.Context) map[string]string {
	places := dal.AllPlace(dal.GetDB())

	placeNameToIdMap := make(map[string]string, len(places))

	for _, place := range places {
		placeNameToIdMap[place.Name] = util.UintToString(place.ID)
	}

	return placeNameToIdMap
}

// GetLocationMapAndPlaceInfo 获取某个place的基础信息以及其下的所有摊位信息
// 先根据placeId拿到其下面的location列表，然后遍历每一个location，组装数据
// 拿到location，去访问其今天的所有order，根据其time-status信息构造结果
// 遍历8-22的每一个时间，从order中如果可以找到，那么根据其订单状态确定这个时间段内该摊位的可预约状态
// 如果找不到，那么根据时间是否是过去式来进行判断
func GetLocationMapAndPlaceInfo(c *gin.Context, placeId uint) map[string]interface{} {
	// 验证区域是否存在
	place := dal.SelectPlace(dal.FilterById(dal.GetDB(), placeId))
	if place == nil {
		log.Printf("[service][place][GetLocationMapAndPlaceInfo] this place is not exist, place_id:%d", placeId)
		panic(constants.NULL_ERROR)
	}

	// 获取区域下的所有摊位
	locationsInThisPlace := GetLocationsByPlaceId(c, place.ID)

	// 返回值
	// "data":{
	//        "placeName":"西邮",
	//        "countLocation":37,
	//        "allLocation":98,
	//        "locationMap":{
	//            "locationId1":{
	//                "introduction":"摊位简介",
	//                "timeStatus"{
	//                    "8":0(可预约)
	//                    "10":1（预约中）
	//                    "12":2（使用中）
	//                    "14":3（不可预约，时间因素，或者当前时间该摊位上有预约单已完成或者过期）
	//                    ......
	//                }
	//            }
	//             "locationId2":{
	//                "introduction":"摊位简介",
	//                "timeStatus"{
	//                    "8":0(可预约)
	//                    "10":1（预约中）
	//                    "12":2（使用中）
	//                    ......
	//                }
	//            }
	//        }
	ans := make(map[string]interface{}, 4)
	// 摊位信息map字段
	locationMap := make(map[string]interface{}, len(locationsInThisPlace))
	ans["location_map"] = locationMap
	ans["all_location"] = len(locationsInThisPlace)
	countAvailableLocation := 0

	for _, location := range locationsInThisPlace {
		// 不可用摊位
		if location.Status == constants.LOCATION_STATUS_UNAVAILABLE {
			continue
		}
		countAvailableLocation++

		// 构造摊位信息
		locationInfo := make(map[string]interface{}, 2)
		locationMap[util.UintToString(location.ID)] = locationInfo
		locationInfo[constants.INTRODUCTION] = location.Introduction

		// 该摊位的时间-可预约信息
		locationStatus := make(map[string]int, 8)
		locationInfo[constants.TIME_STATUS] = locationStatus

		timeToStatusMap := GetAllTodayReserveByLocation(c, location.ID)
		for _, i := range []int{8, 10, 12, 14, 16, 18, 20, 22} {
			s := strconv.Itoa(i)
			if status, isOk := timeToStatusMap[s]; isOk {
				// 当前时段存在预约单
				if status == constants.ORDER_STATUS_TO_BE_USED || status == constants.ORDER_STATUS_IN_USING {
					// 预约单为待使用和使用中状态，摊位可预约状态保持一致
					locationStatus[s] = status
				} else {
					// 预约单为其他状态，已完成和过期，都表示过去式，摊位为不可预约
					locationStatus[s] = constants.LOCATION_USED_STATUS_CAN_NOT_BE_BOOKED
				}
			} else {
				// 不存在预约单，那么通过时间判断可否预约
				nowHour := time.Now().Hour()
				if nowHour < i {
					// 遍历到的时间点是一个将来时间
					locationStatus[s] = constants.LOCATION_USED_STATUS_CAN_BE_BOOKED
				} else {
					// 当前已经过了预约时间了
					locationStatus[s] = constants.LOCATION_USED_STATUS_CAN_NOT_BE_BOOKED
				}
			}
		}
	}
	ans["count_location"] = countAvailableLocation

	return ans
}
