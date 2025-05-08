package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LibraryModel = (*customLibraryModel)(nil)

type (
	// LibraryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLibraryModel.
	LibraryModel interface {
		libraryModel
		withSession(session sqlx.Session) LibraryModel
	}

	customLibraryModel struct {
		*defaultLibraryModel
	}
)

// NewLibraryModel returns a model for the database table.
func NewLibraryModel(conn sqlx.SqlConn) LibraryModel {
	return &customLibraryModel{
		defaultLibraryModel: newLibraryModel(conn),
	}
}

func (m *customLibraryModel) withSession(session sqlx.Session) LibraryModel {
	return NewLibraryModel(sqlx.NewSqlConnFromSession(session))
}
