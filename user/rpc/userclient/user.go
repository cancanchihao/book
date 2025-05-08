package userclient

import (
	"context"

	"book/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	UserExistReq   = user.UserExistReq
	UserExistReply = user.UserExistReply

	User interface {
		IsUserExist(ctx context.Context, in *UserExistReq) (*UserExistReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// IsUserExist 判断用户是否存在
func (m *defaultUser) IsUserExist(ctx context.Context, in *UserExistReq) (*UserExistReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsUserExist(ctx, in)
}
