package model

import (
	"github.com/jinzhu/gorm"
	"street_stall/biz/domain/model/base"
)

// Question 用户反馈问题
type Question struct {
	gorm.Model
	base.Row
	UserId   uint   `gorm:"column:user_id"`  // 提交问题的用户
	Question string `gorm:"column:question"` // 反馈的问题
	Status   int    `gorm:"column:status"`   // 当前处理状态
}
