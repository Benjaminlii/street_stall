package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
)

// GetAdministratorByUsernameAndPassword 根据用户名密码查找管理员
func GetAdministratorByUsernameAndPassword(username string, password string) *model.Administrator {
	db := GetDB()
	db = filterByUsernameAndPassword(db, username, password)
	administrator := selectAdministrator(db)
	return administrator
}

// selectAdministrator 根据db去查询administrator模型
func selectAdministrator(db *gorm.DB) *model.Administrator {
	administrator := &model.Administrator{}
	result := db.First(administrator)
	if err := result.Error; err != nil {
		log.Printf("[service][administrator][selectAdministrator] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return administrator
}
