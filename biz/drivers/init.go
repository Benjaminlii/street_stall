package drivers

import (
	"street_stall/biz/config"
	"sync"
)

var(
	once sync.Once
)

func InitFromConfigOnce() {
	once.Do(func() {
		InitMySQL(&config.AppConfig)
		InitRedis(&config.AppConfig)
	})
}

