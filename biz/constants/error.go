package constants

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		CODE:          e.Code,
		ERROR_MESSAGE: e.ErrorMessage,
		DATA:          dateInterface,
	}
}
