package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/domain/model"
)

// SelectUser 根据db去查询user模型
func SelectUser(db *gorm.DB) *model.User {
	user := &model.User{}
	db.First(user)
	if err := db.Error; err != nil {
		log.Printf("[service][user][SelectUser] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(constants.SYSTEM_ERROR)
		}
	}

	return user
}

// FilterByUsernameAndPassword 通过用户名以及密码
func FilterByUsernameAndPassword(db *gorm.DB, username string, password string) *gorm.DB {
	db = db.Where("username = ?", username)
	db = db.Where("password = ?", password)
	return db
}

// InsertUser 插入一个user对象
func InsertUser(db *gorm.DB, insertUser *model.User) *model.User {
	db.Create(insertUser)
	if err := db.Error; err != nil {
		log.Printf("[service][user][InsertUser] db insert error, err:%s", err)
	}
	return insertUser
}

// FilterByUserId 通过userId过滤
func FilterByUserId(db *gorm.DB, userId uint) *gorm.DB {
	db = db.Where("user_id = ?", userId)
	return db
}
