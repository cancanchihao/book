package svc

import (
	"book/borrow/api/internal/config"
	"book/borrow/model"
	"book/library/rpc/libraryclient"
	"book/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	BorrowSystemModel model.BorrowSystemModel
	UserRpc           userclient.User
	LibraryRpc        libraryclient.Library
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ur := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	lr := libraryclient.NewLibrary(zrpc.MustNewClient(c.LibraryRpc))
	return &ServiceContext{
		Config:            c,
		BorrowSystemModel: model.NewBorrowSystemModel(conn, c.Cache),
		UserRpc:           ur,
		LibraryRpc:        lr,
	}
}
