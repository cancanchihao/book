package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	_         = iota
	Borrowing // 借阅中
	Return    // 已归还
)

var _ BorrowSystemModel = (*customBorrowSystemModel)(nil)

type (
	// BorrowSystemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBorrowSystemModel.
	BorrowSystemModel interface {
		borrowSystemModel

		FindOneByUserAndBookNo(ctx context.Context, userId int64, bookNo int64) (*BorrowSystem, error)
		FindOneByBookNo(ctx context.Context, bookNo int64, status int64) (*BorrowSystem, error)
	}

	customBorrowSystemModel struct {
		*defaultBorrowSystemModel
	}
)

// NewBorrowSystemModel returns a model for the database table.
func NewBorrowSystemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BorrowSystemModel {
	return &customBorrowSystemModel{
		defaultBorrowSystemModel: newBorrowSystemModel(conn, c, opts...),
	}
}
