package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Cache      cache.CacheConf
	LibraryRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
