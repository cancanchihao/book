package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Cache cache.CacheConf
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
}
