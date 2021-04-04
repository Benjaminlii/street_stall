package util

import (
	"log"
	"strconv"
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
