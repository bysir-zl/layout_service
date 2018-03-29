package tool

import (
	"github.com/bysir-zl/bygo/cache"
	"github.com/bysir-zl/bygo/config"
)

var Redis *cache.BRedis

func init() {
	ip := config.GetString("addr", "redis")
	Redis = cache.NewRedis(ip)
}
