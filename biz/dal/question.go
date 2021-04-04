package dal

import (
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/domain/model"
)

// InsertQuestion 插入一个question
func InsertQuestion(db *gorm.DB, insertQuestion *model.Question) *model.Question {
	db.Create(insertQuestion)
	if err := db.Error; err != nil {
		log.Printf("[service][question][InsertQuestion] db insert error, err:%s", err)
	}
	return insertQuestion
}
