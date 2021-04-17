package dao

import (
	"log"
	"street_stall/biz/domain/model"
)

// InsertEvaluation
func InsertEvaluation(insertEvaluation *model.Evaluation) *model.Evaluation {
	db := GetDB()
	db.Create(insertEvaluation)
	if err := db.Error; err != nil {
		log.Printf("[service][evaluation][InsertEvaluation] db insert error, err:%s", err)
		panic(err)
	}
	return insertEvaluation
}
