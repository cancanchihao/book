package logic

import (
	"book/user/api/internal/svc"
	"book/user/api/internal/types"
	"book/user/model"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.UserReply, error) {
	// 忽略逻辑校验
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	switch {
	case err == nil:
		if userInfo.Password != req.Password {
			return nil, errorIncorrectPassword
		}
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
		if err != nil {
			return nil, err
		}

		return &types.UserReply{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
			JwtToken: types.JwtToken{
				AccessToken:  jwtToken,
				AccessExpire: now + accessExpire,
				RefreshAfter: now + accessExpire/2,
			},
		}, nil
	case errors.Is(err, model.ErrNotFound):
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["x-user-id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
