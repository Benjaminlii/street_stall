package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
)

// GetUserById 根据userId获取user
func GetUserById(userId uint) *model.User {
	db := GetDB()
	db = filterById(db, userId)
	user := selectUser(db)
	return user
}

// GetUserByUsernameAndPassword 根据用户名密码查找user
func GetUserByUsernameAndPassword(username string, password string) *model.User {
	db := GetDB()
	db = filterByUsernameAndPassword(db, username, password)
	user := selectUser(db)
	return user
}

// InsertUser 插入一个user对象
func InsertUser(insertUser *model.User) *model.User {
	db := GetDB()
	db.Create(insertUser)
	if err := db.Error; err != nil {
		log.Printf("[service][user][InsertUser] db insert error, err:%s", err)
	}
	return insertUser
}

// selectUser 根据db去查询user模型
func selectUser(db *gorm.DB) *model.User {
	user := &model.User{}
	db.First(user)
	if err := db.Error; err != nil {
		log.Printf("[service][user][selectUser] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return user
}

// filterByUsernameAndPassword 通过用户名以及密码
func filterByUsernameAndPassword(db *gorm.DB, username string, password string) *gorm.DB {
	db = db.Where("username = ?", username)
	db = db.Where("password = ?", password)
	return db
}

// filterByUserId 通过userId过滤
func filterByUserId(db *gorm.DB, userId uint) *gorm.DB {
	db = db.Where("user_id = ?", userId)
	return db
}
