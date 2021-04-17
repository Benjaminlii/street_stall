package dao

import (
	"github.com/jinzhu/gorm"
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

// FindEvaluationsByMerchantId 根据merchantId获取商户评价，按照创建时间降序排列
func FindEvaluationsByMerchantId(merchantId uint) []model.Evaluation {
	db := GetDB()
	db = filterByMerchantId(db, merchantId)
	db = orderByCreatedAt(db, true)

	return findEvaluation(db)
}

// findEvaluation 根据传入的db查询evaluation
func findEvaluation(db *gorm.DB) (evaluations []model.Evaluation) {
	db.Find(&evaluations)
	return
}
