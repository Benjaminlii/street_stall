package util

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/drivers"
	"strings"
	"sync"
	"time"
)

// Lock 加redis分布式锁
func Lock(key string) {
	// 当前方法加锁
	stringLock(key)
	defer stringUnLock(key)

	// 设置自旋次数
	count := 0

	for {
		// 获取锁
		isLock, err := drivers.RedisClient.SetNX(constants.REDIS_LOCK_KEY_PRE+key, constants.REDIS_DEFAULT_VALUE, time.Minute).Result()
		if err != nil {
			log.Printf("[system][lock][Lock] get lock error, err:%s, key:%s", err, key)
			panic(err)
		}
		count++
		if !isLock {
			// 未获取到锁
			// 自旋间隔时间0.1秒
			time.Sleep(time.Millisecond * 100)
		} else {
			// 加锁成功
			log.Printf("[system][lock][Lock] get lock success, key:%s", key)
			break
		}
		// 超出自旋次数则抛异常出去
		if count > 10 {
			log.Printf("[system][lock][Lock] get lock error, spin so mach, key:%s", key)
			panic(errors.LOCK_ERROR)
		}
	}
}

// 释放redis分布式锁
func UnLock(key string) {
	// 当前方法加锁
	stringLock(key)
	defer stringUnLock(key)

	deleteCount, err := drivers.RedisClient.Del(constants.REDIS_LOCK_KEY_PRE + key).Result()
	if err != nil {
		log.Printf("[system][lock][UnLock] free lock error, err:%s, key:%s", err, key)
		panic(err)
	}
	if deleteCount != 1 {
		log.Printf("[system][lock][UnLock] free lock error, delete count:%d", deleteCount)
	}
}

// 监视器锁
var mutex sync.Mutex

// 用于单机key加锁的map
var stringMap = make(map[string]int, 10)

// stringLock 单机对key进行加锁
func stringLock(key string) bool {
	// go的map线程不安全，需要加锁保证线程安全，这个mutex粒度应该足够小了
	mutex.Lock()
	defer mutex.Unlock()

	// 设置自旋次数
	count := 0

	for {
		count++
		_, isOk := stringMap[key]
		if !isOk {
			// 找不到说明可以加单机锁
			stringMap[key] = goroutineId()
			return true
		} else {
			// 未获取到锁
			// 自旋间隔时间0.1秒
			time.Sleep(time.Millisecond * 100)
		}
		// 超出自旋次数则抛异常出去
		if count > 10 {
			log.Printf("[system][lock][stringLock] get string lock error, spin so mach, key:%s", key)
			panic(errors.LOCK_ERROR)
		}
	}
}

// stringUnLock 单机对key进行解锁
func stringUnLock(key string) {
	// go的map线程不安全，需要加锁保证线程安全，这个mutex粒度应该足够小了
	mutex.Lock()
	defer mutex.Unlock()
	delete(stringMap, key)
}

// 获取当前的goroutine id
func goroutineId() int {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[system][lock][goroutineId] panic recover:panic info:%v", err)
		}
	}()
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
