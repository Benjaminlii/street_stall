package dto

type GetEvaluationsDTO struct {
	ID      uint   `json:"id"`      // 评价id
	Star    uint   `json:"star"`    // 评价星级
	Content string `json:"content"` // 评价内容
	Visitor struct {
		UserId       uint   `json:"user_id"`      // 用户id
		Name         string `json:"name"`         // 游客昵称
		Introduction string `json:"introduction"` // 个人简介
	}
}
