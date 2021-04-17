package errors

// 正确
var SUCCESS = &Error{Code: 10000, ErrorMessage: "接口访问成功！"}

// 系统错误
var SYSTEM_ERROR = &Error{Code: 20001, ErrorMessage: "系统错误！"}
var REQUEST_TYPE_ERROR = &Error{Code: 20002, ErrorMessage: "请求参数不合法！"}
var OUTSIDE_ERROR = &Error{Code: 20003, ErrorMessage: "外部系统错误！"}
var OTHER_ERROR = &Error{Code: 20004, ErrorMessage: "其他未知错误:"}
var JSON_ERROR = &Error{Code: 20005, ErrorMessage: "JSON数据格式转化错误！"}
var LOCK_ERROR = &Error{Code: 20006, ErrorMessage: "获取分布式锁失败！"}

// 业务错误
var TOKEN_WRONG_ERROR = &Error{Code: 30001, ErrorMessage: "token无效！"}
var NO_TOKEN_ERROR = &Error{Code: 30002, ErrorMessage: "无token！"}
var AUTHORITY_ERROR = &Error{Code: 30003, ErrorMessage: "无权限！"}
var NULL_ERROR = &Error{Code: 30004, ErrorMessage: "空返回值！"}
var NO_LOGIN_ERROR = &Error{Code: 30005, ErrorMessage: "无登录状态！"}
var LOGIN_FAILD_ERROR = &Error{Code: 30006, ErrorMessage: "登录失败！"}
var ORDER_MERCHANT_ERROR = &Error{Code: 30007, ErrorMessage: "该订单与当前用户不匹配！"}
var ORDER_RESERVE_TIME_ERROR = &Error{Code: 30008, ErrorMessage: "该订单预约时间与当前时间不匹配！"}
