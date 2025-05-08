package model

import (
	"context"
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
		withSession(session sqlx.Session) BorrowSystemModel

		FindOneByUserAndBookNo(ctx context.Context, userId int64, bookNo string) (*BorrowSystem, error)
		FindOneByBookNo(ctx context.Context, bookNo string, status int64) (*BorrowSystem, error)
	}

	customBorrowSystemModel struct {
		*defaultBorrowSystemModel
	}
)

// NewBorrowSystemModel returns a model for the database table.
func NewBorrowSystemModel(conn sqlx.SqlConn) BorrowSystemModel {
	return &customBorrowSystemModel{
		defaultBorrowSystemModel: newBorrowSystemModel(conn),
	}
}

func (m *customBorrowSystemModel) withSession(session sqlx.Session) BorrowSystemModel {
	return NewBorrowSystemModel(sqlx.NewSqlConnFromSession(session))
}
