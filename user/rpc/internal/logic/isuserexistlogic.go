package logic

import (
	"book/user/model"
	"book/user/rpc/user"
	"context"
	"errors"

	"book/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserExistLogic {
	return &IsUserExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IsUserExist 判断用户是否存在
func (l *IsUserExistLogic) IsUserExist(in *user.UserExistReq) (*user.UserExistReply, error) {
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	switch {
	case err == nil:
		return &user.UserExistReply{Exists: true}, nil
	case errors.Is(err, model.ErrNotFound):
		return &user.UserExistReply{}, nil
	default:
		return nil, err
	}

}
