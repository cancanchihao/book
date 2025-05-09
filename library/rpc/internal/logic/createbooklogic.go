package logic

import (
	"book/library/model"
	"book/library/rpc/internal/svc"
	"book/library/rpc/library"
	"book/shared"
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type CreateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookLogic {
	return &CreateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBookLogic) CreateBook(in *library.CreateBookReq) (*library.CreateBookReply, error) {
	var publishDate sql.NullTime
	if in.PublishDate != "" {
		t, err := time.Parse(timeFormat, in.PublishDate)
		if err != nil {
			return nil, fmt.Errorf("invalid publish date format: %v", err)
		}
		publishDate = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	_, err := l.svcCtx.LibraryModel.Insert(l.ctx, &model.Library{
		Name:        in.Name,
		Author:      in.Author,
		PublishDate: publishDate,
	})

	if err == nil {
		return &library.CreateBookReply{
			Success: err == nil,
		}, nil
	} else {
		return nil, shared.NewGRPCErrorFromError(err)
	}
}
