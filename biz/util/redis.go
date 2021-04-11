package util

import (
	"encoding/json"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
	"street_stall/biz/drivers"
	"time"
)
import "github.com/satori/go.uuid"

// AddUserToken 向redis中添加某个用户的token，有效时间为3天
func AddUserToken(user *model.User) (token string) {
	// 生成该用户的token
	token = uuid.NewV4().String()
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("[system][redis] json marshal error, err:%s", err)
		panic(errors.JSON_ERROR)
	}
	drivers.RedisClient.Set(constants.REDIS_USER_TOKEN_PRE+token, userJson, time.Hour*24*3)
	return token
}
