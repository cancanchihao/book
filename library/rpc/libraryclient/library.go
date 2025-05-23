// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: library.proto

package libraryclient

import (
	"context"

	"book/library/rpc/library"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateBookReply = library.CreateBookReply
	CreateBookReq   = library.CreateBookReq
	FindBookReply   = library.FindBookReply
	FindBookReq     = library.FindBookReq

	Library interface {
		FindBookByName(ctx context.Context, in *FindBookReq, opts ...grpc.CallOption) (*FindBookReply, error)
		CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookReply, error)
	}

	defaultLibrary struct {
		cli zrpc.Client
	}
)

func NewLibrary(cli zrpc.Client) Library {
	return &defaultLibrary{
		cli: cli,
	}
}

// FindBookByName 通过书籍名称查找书籍
func (m *defaultLibrary) FindBookByName(ctx context.Context, in *FindBookReq, opts ...grpc.CallOption) (*FindBookReply, error) {
	client := library.NewLibraryClient(m.cli.Conn())
	return client.FindBookByName(ctx, in, opts...)
}

func (m *defaultLibrary) CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookReply, error) {
	client := library.NewLibraryClient(m.cli.Conn())
	return client.CreateBook(ctx, in, opts...)
}
