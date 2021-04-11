package dto

type GetOrderDTO struct {
	OrderId  uint  `json:"order_id"`  // 预约单id
	CreateAt int64 `json:"create_at"` // 预约单创建时间
	Status   int   `json:"status"`    // 当前订单状态  0（待使用）/1（使用中）/2（已完成）/3（过期）
	Location struct {
		Place struct {
			Name string `json:"name"` // 摊位所属区域名称
		} `json:"place"` // 摊位所在区域
		Number       int    `json:"number"`       // 摊位编号(在某区域内)
		Introduction string `json:"introduction"` // 摊位简介
	} `json:"location"` // 该订单所预定摊位信息
	ReserveTime uint   `json:"reserve_time"` // 预约时间
	Remark      string `json:"remark"`       // 预约单备注
}
