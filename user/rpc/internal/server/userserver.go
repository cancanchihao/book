package server

import (
	"context"

	"book/user/rpc/internal/logic"
	"book/user/rpc/internal/svc"
	"book/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// IsUserExist 判断用户是否存在
func (s *UserServer) IsUserExist(ctx context.Context, in *user.UserExistReq) (*user.UserExistReply, error) {
	l := logic.NewIsUserExistLogic(ctx, s.svcCtx)
	return l.IsUserExist(in)
}
