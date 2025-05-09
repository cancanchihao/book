package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var cacheBookBorrowSystemBookNoStatusPrefix = "cache:book:borrowSystem:bookNo:status:"

func (m *customBorrowSystemModel) FindOneByUserAndBookNo(ctx context.Context, userId int64, bookNo int64) (*BorrowSystem, error) {
	bookBorrowSystemUserIdBookNoKey := fmt.Sprintf("%s%v:%v", cacheBookBorrowSystemUserIdBookNoPrefix, userId, bookNo)
	var resp BorrowSystem

	err := m.QueryRowCtx(ctx, &resp, bookBorrowSystemUserIdBookNoKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf(
			"SELECT %s FROM %s WHERE `user_id` = ? AND `book_no` = ? LIMIT 1",
			borrowSystemRows, m.tableName(),
		)
		return conn.QueryRowCtx(ctx, v, query, userId, bookNo)
	})
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customBorrowSystemModel) FindOneByBookNo(ctx context.Context, bookNo int64, status int64) (*BorrowSystem, error) {
	bookBorrowSystemUserIdBookNoStatusKey := fmt.Sprintf("%s%v:%v", cacheBookBorrowSystemBookNoStatusPrefix, bookNo, status)
	var resp BorrowSystem

	err := m.QueryRowCtx(ctx, &resp, bookBorrowSystemUserIdBookNoStatusKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf(
			"SELECT %s FROM %s WHERE `book_no` = ? AND `status` = ? LIMIT 1",
			borrowSystemRows, m.tableName(),
		)
		return conn.QueryRowCtx(ctx, v, query, bookNo, status)
	})
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
