package service

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/dal"
	"street_stall/biz/domain/model"
)

// GetLocationsByPlaceId 获取某个place（区域）下的所有Location（摊位）
func GetLocationsByPlaceId(c *gin.Context, placeId uint) []model.Location {
	db := dal.GetDB()
	db = dal.FilterByPlaceId(db, placeId)
	locations := dal.FindLocation(db)
	return locations
}
