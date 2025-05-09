package logic

import (
	"book/library/model"
	"book/shared"
	"context"
	"errors"

	"book/library/rpc/internal/svc"
	"book/library/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	timeFormat = "2006-01-02"
)

type FindBookByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBookByNameLogic {
	return &FindBookByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindBookByNameLogic) FindBookByName(in *library.FindBookReq) (*library.FindBookReply, error) {
	book, err := l.svcCtx.LibraryModel.FindOneByName(l.ctx, in.Name)
	switch {
	case err == nil:
		var publishDate string
		if book.PublishDate.Valid {
			publishDate = book.PublishDate.Time.Format(timeFormat)
		}
		return &library.FindBookReply{
			No:          book.Id,
			Name:        book.Name,
			Author:      book.Author,
			PublishDate: publishDate,
		}, nil
	case errors.Is(err, model.ErrNotFound):
		return nil, shared.NewGRPCNotFound()
	default:
		return nil, shared.NewGRPCErrorFromError(err)
	}

}
