package logic

import (
	"book/user/model"
	"context"
	"errors"
	"strconv"

	"book/user/api/internal/svc"
	"book/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(userId string) (*types.UserReply, error) {
	userInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userInt)
	switch {
	case err == nil:
		return &types.UserReply{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
		}, nil
	case errors.Is(err, model.ErrNotFound):
		return nil, errorUserNotFound
	default:
		return nil, err
	}
}
