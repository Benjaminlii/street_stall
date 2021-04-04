package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"street_stall/biz/constants"
)

func SetResponse(c *gin.Context) {

	// data字段转化
	data, err := c.Get("data")
	dataInterface := new(interface{})
	if !err {
		// 接口中未定义data字段
		log.Print("[system][error] response has not data field.")
		dataInterface = nil
	} else {
		dataInterface = &data
	}

	// 发生错误
	resp := constants.SUCCESS.ChangeToResp(dataInterface)
	if err := recover(); err != nil {
		// 已定义错误
		if myError, isOk := err.(constants.Error); isOk {
			resp = myError.ChangeToResp(dataInterface)
		} else {
			resp = constants.OTHER_ERROR.ChangeToResp(dataInterface)
			resp[constants.ERROR_MESSAGE] = fmt.Sprintf("%s%s", constants.OTHER_ERROR.ErrorMessage, err)
		}
	}

	c.JSON(http.StatusOK, resp)
}
