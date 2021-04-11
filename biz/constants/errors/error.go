package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"street_stall/biz/constants"
)

// Error 自定义错误类型
type Error struct {
	Code         int
	ErrorMessage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d, errorMessage:%s", e.Code, e.ErrorMessage)
}

// ChangeToResp 根据错误和传入的data生成标准响应
func (e *Error) ChangeToResp(dateInterface interface{}) gin.H {
	return gin.H{
		constants.CODE:          e.Code,
		constants.ERROR_MESSAGE: e.ErrorMessage,
		constants.DATA:          dateInterface,
	}
}
