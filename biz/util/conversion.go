package util

import (
	"log"
	"strconv"
	"street_stall/biz/constants"
)

// StringToUInt string类型转换为uint类型，十进制
func StringToUInt(str string) uint {
	gotUint64, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		log.Printf("[system][util][StringToUInt] change error, err:%s", err)
		panic(err)
	}
	return uint(gotUint64)
}

// UintToString uint类型转换为string类型
func UintToString(i uint) string {
	str := strconv.Itoa(int(i))
	return str
}

// UintToCategoryString 商户分类uint转为内容文字
func UintToCategoryString(i uint) string {
	ans := "无效商户分类"
	switch i {
	case constants.CATEGORY_CULTURE:
		ans = constants.CATEGORY_CULTURE_STR
	case constants.CATEGORY_FOOD:
		ans = constants.CATEGORY_FOOD_STR
	case constants.CATEGORY_LOCATION:
		ans = constants.CATEGORY_LOCATION_STR
	}
	return ans
}
