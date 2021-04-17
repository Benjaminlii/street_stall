package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"street_stall/biz/dao"
	"street_stall/biz/domain/model"
	"street_stall/biz/util"
)

// DoEvaluation 游客对商户进行评价
func DoEvaluation(c *gin.Context, merchantId uint, start uint, content string) {
	// 校验merchant
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
