package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// SubmitQuestion 用户问题反馈
func SubmitQuestion(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][question][SubmitQuestion] request type error, err:%s", err)
		panic(err)
	}
	question, haveQuestion := param["question"]
	if !haveQuestion {
		log.Print("[service][question][SubmitQuestion] there is no question")
		panic(constants.REQUEST_TYPE_ERROR)
	}

	insertedQuestion := service.SaveQuestionByCurrentUser(c, question)

	log.Printf("[service][question][SubmitQuestion] submit question success, userId:%s, question:%s",
		insertedQuestion.UserId, insertedQuestion.Question)

	// 设置请求响应
	respMap := map[string]interface{}{}

	c.Set(constants.DATA, respMap)
}
