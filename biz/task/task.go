package task

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/drivers"
	"sync"
	"time"
)

var (
	once sync.Once
)

func InitTask() {
	once.Do(func() {
		c := cron.New()
		err := c.AddFunc("0 0 0/2 * * ? *", taskDeleteCurrentMerchantEveryTwoHour)
		if err != nil {
			log.Printf("[system][task][InitTask] add task error, task name is:taskDeleteCurrentMerchantEveryTwoHour")
		}
		err = c.AddFunc("0 0 0/2 * * ? *", taskUpdateOrderEveryTwoHour)
		if err != nil {
			log.Printf("[system][task][InitTask] add task error, task name is:taskUpdateOrderEveryTwoHour")
		}
		c.Start()

		log.Println("[system][task][InitTask] init task success!")
	})
}

// taskDeleteCurrentMerchantEveryTwoHour,redis活跃商户清空定时任务,每两个小时将redis中所有当前活跃（正在摆摊）的商户set删除
func taskDeleteCurrentMerchantEveryTwoHour() {
	time1 := time.Now()
	log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] task start, current time:%s", time1)

	// 获取当前的时间段值（8/10/12/14等等），如果不在可选范围内，结束任务
	nowHour := time.Now().Hour()
	rightHour := []int{10, 12, 14, 16, 18, 20, 22, 0}
	flag := false
	for _, hour := range rightHour {
		if nowHour == hour {
			flag = true
		}
	}
	if !flag {
		log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] current time is not a right time to run this task.")
		return
	}

	// 获取所有的区域
	places := dao.AllPlace()

	// 拼凑得到所有的redis key
	redisKeys := make([]string, 0)
	for _, place := range places {
		redisKeys = append(redisKeys, fmt.Sprintf("%s%d", constants.REDIS_CURRENT_ACTIVE_MERCHANT_PRE, place.ID))
	}

	// 删除
	for _, redisKey := range redisKeys {
		delCount, err := drivers.RedisClient.Del(redisKey).Result()
		if delCount == 1 {
			log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] delete key success, key:%s", redisKey)
		} else if delCount == 0 {
			log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] delete key fail, key:%s", redisKey)
		}
		if err != nil {
			log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] delete key error, err:%s", err)
			panic(errors.SYSTEM_ERROR)
		}
	}
	time2 := time.Now()
	usedSecond := time2.Sub(time1)
	log.Printf("[system][task][taskDeleteCurrentMerchantEveryTwoHour] task end, used cesond:%s", usedSecond)
}

// taskUpdateOrderEveryTwoHour,order更新状态定时任务（状态从使用中到完成）
// 每两个小时将redis中正在摆摊的预约项更新为已完成（即 使用中 -> 到时间 -> 已完成）
// 将未打卡的预约项更新为过期（即 待使用 -> 到时间 -> 过期）
func taskUpdateOrderEveryTwoHour() {
	time1 := time.Now()
	log.Printf("[system][task][taskUpdateOrderEveryTwoHour] task start, current time:%s", time1)

	// 获取当前的时间段值（8/10/12/14等等），如果不在可选范围内，结束任务
	nowHour := time.Now().Hour()
	rightHourMapping := map[int]uint{
		0:  22,
		10: 8,
		12: 10,
		14: 12,
		16: 14,
		18: 16,
		20: 18,
		22: 20,
	}
	rightHour, isOk := rightHourMapping[nowHour]
	if !isOk {
		log.Printf("[system][task][taskUpdateOrderEveryTwoHour] current time is not a right time to run this task.")
	}

	// 过滤得到今天，当前时间段，并且正在被使用的订单
	ifYesterday := false
	if rightHour == 22 {
		ifYesterday = true
	}
	orders := dao.GetTodayOrderByStatusAndReserveTime(constants.ORDER_STATUS_IN_USING, rightHour, ifYesterday)
	// 更新
	for _, order := range orders {
		status := constants.ORDER_STATUS_FINISHED
		if order.Status == constants.ORDER_STATUS_IN_USING {
			// 使用中 -> 到时间 -> 已完成
			status = constants.ORDER_STATUS_FINISHED
		} else if order.Status == constants.ORDER_STATUS_TO_BE_USED {
			// 待使用 -> 到时间 -> 过期
			status = constants.ORDER_STATUS_EXPIRED
		}
		order.Status = status
		dao.SaveOrder(&order)
	}

	time2 := time.Now()
	usedSecond := time2.Sub(time1)
	log.Printf("[system][task][taskUpdateOrderEveryTwoHour] task end, used cesond:%s", usedSecond)
}
