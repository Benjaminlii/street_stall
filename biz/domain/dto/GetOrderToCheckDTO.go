package dto

type GetOrderToCheckDTO struct {
	OrderId       uint   `json:"order_id"`        // 预约单id
	MerchantName  string `json:"merchant_name"`   // 提交预约的商户名
	PlaceName     string `json:"place_name"`      // 预约的地区名
	NumberOfPlace int    `json:"number_of_place"` // 所预约摊位的偏移量
	Time          uint   `json:"time"`            // 预约的时间段
	PostTime      int64  `json:"post_time"`       // 提交预约的时间点
	Remark        string `json:"remark"`          // 预约信息
}
