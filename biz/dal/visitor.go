package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// InsertVisitor 插入一个visitor对象
func InsertVisitor(db *gorm.DB, insertVisitor *model.Visitor) *model.Visitor {
	db.Create(insertVisitor)
	if db.Error != nil {
		return nil
	}
	return insertVisitor
}

// SelectVisitor 查询visitor
func SelectVisitor(db *gorm.DB) *model.Visitor{
	visitor := &model.Visitor{}
	db = db.First(visitor)
	if err := db.Error; err != nil {
		log.Printf("[service][visitor][SelectVisitor] db select error, err:%s", err)
	}
	return visitor
}

// SaveVisitor 更新并覆盖merchant
func SaveVisitor(db *gorm.DB, visitor *model.Visitor) {
	db.Save(visitor)
}