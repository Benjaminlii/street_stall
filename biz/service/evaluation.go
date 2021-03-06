package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"street_stall/biz/dao"
	"street_stall/biz/domain/dto"
	"street_stall/biz/domain/model"
	"street_stall/biz/util"
)

// DoEvaluation 游客对商户进行评价
func DoEvaluation(c *gin.Context, merchantId uint, start uint, content string) {
	// 校验merchantId
	merchant := dao.GetMerchantById(merchantId)
	// 插入一条评价记录
	visitor := GetVisitorByCurrentUser(c)
	evaluation := &model.Evaluation{
		Star:       start,
		Content:    content,
		MerchantId: merchantId,
		VisitorId:  visitor.ID,
	}
	evaluation = dao.InsertEvaluation(evaluation)

	// 更新商户的基础信息，评价数目和累积星数
	lockKey := fmt.Sprintf("update_merchant_%d", merchant.ID)
	util.Lock(lockKey)
	defer util.UnLock(lockKey)

	merchant.CommentCount++
	merchant.StarSum += start

	dao.SaveMerchant(merchant)
}

// GetEvaluationsByMerchantId 根据商户id获取的评价信息
func GetEvaluationsByMerchantId(c *gin.Context, merchantId uint) []dto.GetEvaluationsDTO {
	// 校验merchantId
	merchant := dao.GetMerchantById(merchantId)

	// 根据merchantId获取evaluations
	evaluations := dao.FindEvaluationsByMerchantId(merchant.ID)
	ans := make([]dto.GetEvaluationsDTO, 0)

	for _, evaluation := range evaluations {
		visitor := dao.GetVisitorById(evaluation.VisitorId)
		getEvaluationsDTO := dto.GetEvaluationsDTO{
			ID:      evaluation.ID,
			Star:    evaluation.Star,
			Content: evaluation.Content,
			Visitor: struct {
				UserId         uint   `json:"user_id"`         // 用户id
				Name           string `json:"name"`            // 游客昵称
				Introduction   string `json:"introduction"`    // 个人简介
				EvaluationDate int64  `json:"evaluation_date"` // 评价时间
			}{UserId: visitor.ID,
				Name:           visitor.Name,
				Introduction:   visitor.Introduction,
				EvaluationDate: evaluation.CreatedAt.UnixNano() / 1e6},
		}
		ans = append(ans, getEvaluationsDTO)
	}

	return ans
}
