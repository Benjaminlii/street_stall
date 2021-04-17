package main

import (
	"github.com/gin-gonic/gin"
	"street_stall/biz/config"
	"street_stall/biz/drivers"
	"street_stall/biz/task"
)

// 主函数
func main() {
	r := gin.Default()

	config.InitConfig("conf/config.yml")
	drivers.InitFromConfigOnce()
	task.InitTask()

	register(r)
	r.Run(":8585")
}
