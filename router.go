package main

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/handler"
	"street_stall/biz/middleware"
)

func register(r *gin.Engine) {
	streetStall := r.Group("/api/street_stall")

	// 用户模块
	user := streetStall.Group("/user")
	{
		user.POST("/sign_in", handler.SignIn)
		user.POST("/sign_up", handler.SignUp)

		user.Use(middleware.CheckUserLoginMiddleware())
		user.POST("/get_merchant", handler.GetMerchant)
		user.POST("/update_merchant", handler.UpdateMerchant)
		user.POST("/get_visitor", handler.GetVisitor)
		user.POST("/update_visitor", handler.UpdateVisitor)
		user.POST("/sign_out", handler.SignOut)
	}

	streetStall.Use(middleware.CheckUserLoginMiddleware())

	// 问题反馈模块
	question := streetStall.Group("/question")
	{
		question.POST("/submit_question", handler.SubmitQuestion)
	}

	// 区域管理模块
	place := streetStall.Group("/place")
	{
		place.POST("/get_place_name_to_id_map", handler.GetPlaceNameToIdMap)
		place.POST("/get_location_map", handler.GetLocationMap)
	}

	// 摊位管理模块
	location := streetStall.Group("/location")
	{
		location.POST("/reserve", handler.Reserve)
	}

	ping := streetStall.Group("/ping")
	{
		ping.POST("/ping", handler.Ping)
	}

}
