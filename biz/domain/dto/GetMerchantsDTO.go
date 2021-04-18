package dto

type GetMerchantsDTO struct {
	MerchantId   uint    `json:"merchant_id"`   // 商户id
	MerchantName string  `json:"merchant_name"` // 商家名称
	Stars        float64 `json:"stars"`         // 商户星级
	Introduction string  `json:"introduction"`  // 商户简介
	Location     struct {
		LocationId    uint `json:"location_id"`     // 商户当前所在摊位id
		NumberOfPlace int  `json:"number_of_place"` // 该摊位的区域偏移量
	}
}
