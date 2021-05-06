package constants

// 用户身份
const (
	USERIDENTITY_MERCHANT = 1 // 商家
	USERIDENTITY_VISITER  = 2 // 游客
	USERIDENTITY_MASTER   = 3 // 政府管理员
)

// 商户分类
const (
	CATEGORY_CULTURE      = 1 // 文化类
	CATEGORY_CULTURE_STR  = "文化类"
	CATEGORY_FOOD         = 2 // 美食类
	CATEGORY_FOOD_STR     = "美食类"
	CATEGORY_LOCATION     = 3 // 地域特色类
	CATEGORY_LOCATION_STR = "地域特色类"
)

// 订单状态
const (
	ORDER_STATUS_TO_BE_USED     = 1 // 待使用
	ORDER_STATUS_IN_USING       = 2 // 使用中
	ORDER_STATUS_FINISHED       = 3 // 已完成
	ORDER_STATUS_EXPIRED        = 4 // 过期
	ORDER_STATUS_CHECK_FINISHED = 5 // 审核成功待使用
)

// 摊位的可预约状态
const (
	LOCATION_USED_STATUS_CAN_BE_BOOKED     = 0 // 可预约
	LOCATION_USED_STATUS_BE_BOOKED         = 1 // 已预约
	LOCATION_USED_STATUS_IN_USING          = 2 // 使用中
	LOCATION_USED_STATUS_CAN_NOT_BE_BOOKED = 3 // 不可预约
)

// 问题反馈的处理状态
const (
	QUESTION_STATUS_START = 1 // 未处理
	QUESTION_STATUS_DONE  = 3 // 处理结束
)

// 摊位的可使用状态
const (
	LOCATION_STATUS_AVAILABLE   = 1 // 可用
	LOCATION_STATUS_UNAVAILABLE = 2 // 不可用
)

// 系统字段
const (
	TOKEN         = "token"         // 登录令牌
	CODE          = "code"          // 标准响应中的状态码
	ERROR_MESSAGE = "error_message" // 标准响应中的错误信息
	DATA          = "data"          // 标准响应中的数据域
	CURRENT_USER  = "current_user"  // 标准响应中的数据域

)

// 业务字段
const (
	INTRODUCTION = "introduction" // 摊位的介绍
	TIME_STATUS  = "time_status"  // 摊位的时间以及可否预约状态
	AREA         = "area"         // 摊位面积

)

// Redis相关
const (
	REDIS_USER_TOKEN_PRE              = "street_stall_user_token_"              // 当前登录的用户在redis中存储有过期时间键的key前缀
	REDIS_CURRENT_ACTIVE_MERCHANT_PRE = "street_stall_current_active_merchant_" // 当前活跃商家id-摊位id的hash（预约并打卡）
	REDIS_LOCK_KEY_PRE                = "street_stall_redis_lock_key_"          // redis分布式锁key
	REDIS_DEFAULT_VALUE               = "0"                                     // redis无需存储value的value值
	REDIS_MANAGER_TOKEN_PRE           = "street_stall_manager_token_"           // 当前登录的管理员在redis中存储有过期时间键的key前缀
)
