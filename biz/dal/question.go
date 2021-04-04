package dal

import (
	"log"
	"street_stall/biz/domain/model"
)

// InsertQuestion 插入一个question
func InsertQuestion(insertQuestion *model.Question) *model.Question {
	db := GetDB()
	db.Create(insertQuestion)
	if err := db.Error; err != nil {
		log.Printf("[service][question][InsertQuestion] db insert error, err:%s", err)
	}
	return insertQuestion
}
