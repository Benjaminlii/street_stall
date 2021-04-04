package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/dal"
	"street_stall/biz/drivers"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// SignIn 用户登录
func SignIn(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][SignIn] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	if !(haveUsername && havePassword) {
		log.Printf("[service][user][SignIn] request type error, err:%s", err)
		panic(constants.REQUEST_TYPE_ERROR)
	}

	// 校验用户信息
	user := service.SelectUser(username, password)
	if user == nil{
		panic(constants.LOGIN_FAILD_ERROR)
	}

	// 生成并添加token到redis，存储user的json
	token := util.AddUserToken(user)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap[constants.TOKEN] = token
	respMap["login_flag"] = user.UserIdentity
	c.Set(constants.DATA, respMap)
}

// SignUp 用户注册
func SignUp(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][SignUp] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	name, haveName := param["name"]
	userIdentityStr, haveUserIdentity := param["userIdentity"]
	categoryStr, haveCategory := param["category"]
	if !(haveUsername && havePassword && haveName && haveUserIdentity && haveCategory) {
		log.Print("[service][user][SignUp] has nil in username, password, name, userIdentity and category")
		panic(constants.REQUEST_TYPE_ERROR)
	}
	userIdentity := util.StringToUInt(userIdentityStr)
	category := util.StringToUInt(categoryStr)

	// 注册
	user := service.SignUp(username, password, name, userIdentity, category)


	// 生成并添加token到redis，存储user的json
	token := util.AddUserToken(user)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap[constants.TOKEN] = token
	respMap["login_flag"] = user.UserIdentity
	c.Set(constants.DATA, respMap)
}

// UpdateMerchant 商户信息维护
func UpdateMerchant(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][UpdateMerchant] request type error, err:%s", err)
		panic(err)
	}
	name, haveName := param["name"]
	categoryStr, haveCategory := param["category"]
	introduction, haveIntroduction := param["info"]
	if !(haveName && haveCategory && haveIntroduction) {
		log.Print("[service][user][UpdateMerchant] has nil in name, category and introduction")
		panic(constants.REQUEST_TYPE_ERROR)
	}
	category := util.StringToUInt(categoryStr)

	// 更新商户信息
	merchant := service.UpdateMerchantByUserId(c, name, category, introduction)


	// 设置请求响应
	user := dal.GetUserById(merchant.UserId)
	if user == nil {
		panic(constants.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         merchant.Name,
		"category":     merchant.Category,
		"introduction": merchant.Introduction,
		"user": map[string]interface{}{
			"username":     user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// UpdateVisitor 游客信息维护
func UpdateVisitor(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][UpdateVisitor] request type error, err:%s", err)
		panic(err)
	}
	name, haveName := param["name"]
	introduction, haveIntroduction := param["info"]
	if !(haveName && haveIntroduction) {
		log.Print("[service][user][UpdateVisitor] has nil in name and introduction")
		panic(constants.REQUEST_TYPE_ERROR)
	}

	// 更新游客信息
	merchant := service.UpdateVisitorByUserId(c, name, introduction)

	// 设置请求响应
	user := dal.GetUserById(merchant.UserId)
	if user == nil{
		panic(constants.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         merchant.Name,
		"introduction": merchant.Introduction,
		"user": map[string]interface{}{
			"username":     user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// GetVisitor 得到当前游客信息
func GetVisitor(c *gin.Context) {
	defer util.SetResponse(c)

	visitor := service.GetVisitorByCurrentUser(c)

	// 设置请求响应
	user := dal.GetUserById(visitor.UserId)
	if user == nil{
		panic(constants.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         visitor.Name,
		"introduction": visitor.Introduction,
		"user": map[string]interface{}{
			"username":     user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// GetMerchant 得到当前商户信息
func GetMerchant(c *gin.Context) {
	defer util.SetResponse(c)

	merchant := service.GetMerchantByCurrentUser(c)

	// 设置请求响应
	user := dal.GetUserById(merchant.UserId)
	if user == nil{
		panic(constants.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         merchant.Name,
		"introduction": merchant.Introduction,
		"user": map[string]interface{}{
			"username":     user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// SignOut 登出
func SignOut(c *gin.Context) {
	defer util.SetResponse(c)

	// 从headers获取token
	token := c.Request.Header["Token"][0]
	if token == "" {
		log.Printf("[service][user][SignOut] no token")
		panic(constants.NO_TOKEN_ERROR)
	}
	deleteCount, err := drivers.RedisClient.Del(constants.REDIS_USER_TOKEN_PRE + token).Result()
	if err != nil {
		log.Printf("[service][user][LogOut] delete redis key error, err:%s", err)
		panic(err)
	}
	if deleteCount != 1 {
		log.Printf("[service][user][SignOut] delete redis count is 0")
		panic(constants.SYSTEM_ERROR)
	}
	// 设置请求响应
	respMap := map[string]interface{}{}

	c.Set(constants.DATA, respMap)
}
