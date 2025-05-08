package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func (m *customBorrowSystemModel) FindOneByUserAndBookNo(ctx context.Context, userId int64, bookNo string) (*BorrowSystem, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `book_no` = ? limit 1", borrowSystemRows, m.tableName())
	var resp BorrowSystem
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, bookNo)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customBorrowSystemModel) FindOneByBookNo(ctx context.Context, bookNo string, status int64) (*BorrowSystem, error) {
	query := fmt.Sprintf("select %s from %s where `book_no` = ? and `status` = ? limit 1", borrowSystemRows, m.tableName())
	var resp BorrowSystem
	err := m.conn.QueryRowCtx(ctx, &resp, query, bookNo, status)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
