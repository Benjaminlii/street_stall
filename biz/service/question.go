package service

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/constants"
	"street_stall/biz/dal"
	"street_stall/biz/domain/model"
	"street_stall/biz/util"
)

// SaveQuestionByCurrentUser 以当前登录的用户反馈一条问题
func SaveQuestionByCurrentUser(c *gin.Context, question string) *model.Question {
	currentUser := util.GetCurrentUser(c)

	insertQuestion := &model.Question{
		UserId:   currentUser.ID,
		Question: question,
		Status:   constants.QUESTION_STATUS_START,
	}

	q := dal.InsertQuestion(dal.GetDB(), insertQuestion)
	return q
}
