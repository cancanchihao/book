package logic

import (
	"book/borrow/api/internal/svc"
	"book/borrow/api/internal/types"
	"book/borrow/model"
	"book/library/rpc/library"
	"book/shared"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnLogic {
	return ReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnLogic) Return(userId string, req types.ReturnReq) error {
	userInt, err := strconv.ParseInt(fmt.Sprintf("%v", userId), 10, 64)
	if err != nil {
		return err
	}

	book, err := l.svcCtx.LibraryRpc.FindBookByName(l.ctx, &library.FindBookReq{Name: req.BookName})
	if err != nil { // code error
		if shared.IsGRPCNotFound(err) {
			return errBookNotFound
		}
		return err
	}

	info, err := l.svcCtx.BorrowSystemModel.FindOneByUserAndBookNo(l.ctx, userInt, book.No)
	switch {
	case err == nil:
		if info.Status == model.Return {
			return errBookReturn
		}
		info.ReturnDate = time.Now().Unix()
		info.Status = model.Return
		err = l.svcCtx.BorrowSystemModel.Update(l.ctx, info)
		return err
	case errors.Is(err, model.ErrNotFound):
		return errUserReturn
	default:
		return err
	}
}
