package libraryclient

import (
	"context"

	"book/library/rpc/library"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	FindBookReq   = library.FindBookReq
	FindBookReply = library.FindBookReply

	Library interface {
		FindBookByName(ctx context.Context, in *FindBookReq) (*FindBookReply, error)
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
func (m *defaultLibrary) FindBookByName(ctx context.Context, in *FindBookReq) (*FindBookReply, error) {
	client := library.NewLibraryClient(m.cli.Conn())
	return client.FindBookByName(ctx, in)
}
