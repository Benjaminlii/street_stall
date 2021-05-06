package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// GetVisitorByUserId 根据userId获取对应的visitor
func GetVisitorByUserId(userId uint) *model.Visitor {
	db := GetDB()
	db = filterByUserId(db, userId)
	visitor := selectVisitor(db)
	return visitor
}

// GetVisitorById 根据Id获取对应的visitor
func GetVisitorById(id uint) *model.Visitor {
	db := GetDB()
	db = filterById(db, id)
	visitor := selectVisitor(db)
	return visitor
}

// InsertVisitor 插入一个visitor对象
func InsertVisitor(insertVisitor *model.Visitor) *model.Visitor {
	db := GetDB()
	db = db.Create(insertVisitor)
	if err := db.Error; err != nil {
		log.Printf("[service][visitor][InsertVisitor] db insert error, err:%s", err)
		panic(err)
	}
	return insertVisitor
}

// SaveVisitor 更新并覆盖merchant
func SaveVisitor(visitor *model.Visitor) {
	db := GetDB()
	db.Save(visitor)
}

// selectVisitor 查询visitor
func selectVisitor(db *gorm.DB) *model.Visitor {
	visitor := &model.Visitor{}
	result := db.First(visitor)
	if err := result.Error; err != nil {
		log.Printf("[service][visitor][selectVisitor] db select error, err:%s", err)
	}
	return visitor
}
